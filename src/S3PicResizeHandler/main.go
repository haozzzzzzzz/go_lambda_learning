package main

import (
	"bytes"
	"context"
	"fmt"
	"image/jpeg"
	"io"

	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/nfnt/resize"
	"github.com/sirupsen/logrus"
)

const (
	Region = "us-east-1"
)

// 下载原图
func DownloadOriginalImage(writer io.WriterAt, bucket, fileName string) (err error) {
	ses, err := session.NewSession(&aws.Config{
		Region: aws.String(Region),
	})
	downloader := s3manager.NewDownloader(ses)
	numBytes, err := downloader.Download(writer, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileName),
	})
	if nil != err {
		fmt.Errorf("Unable to download %q from %q\n", fileName, bucket)
		log.Fatal(err)
		return
	}

	fmt.Println("Downloaded", fileName, numBytes, "bytes")
	return
}

// 上传缩略图
func UploadThumbnailImage(reader io.Reader, bucket, fileName string) (err error) {
	ses, err := session.NewSession(&aws.Config{
		Region: aws.String(Region),
	})

	uploader := s3manager.NewUploader(ses)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(fileName),
		Body:   reader,
	})
	if nil != err {
		fmt.Errorf("Unable to upload %q to %q\n", fileName, bucket)
		log.Fatal(err)
		return
	}
	fmt.Printf("Successfully uploaded %q to %q\n", fileName, bucket)
	return
}

// 生成缩略图
func ResizeImage(reader io.Reader, writer io.Writer, width uint) (err error) {
	image, err := jpeg.Decode(reader)
	if nil != err {
		log.Fatal(err)
		return
	}
	m := resize.Resize(width, 0, image, resize.Lanczos3)
	if nil != err {
		log.Fatal(err)
		return
	}

	err = jpeg.Encode(writer, m, nil)
	if nil != err {
		log.Fatal(err)
		return
	}

	fmt.Println("resize sucessfully")
	return
}

type OffsetByteBuffer struct {
	bytes.Buffer
}

func (m *OffsetByteBuffer) WriteAt(p []byte, off int64) (n int, err error) {
	n, err = m.Write(p)
	if nil != err {
		logrus.Warnf("write bytes failed. %s", err)
		return
	}
	return
}

func NewOffsetByteBuffer(buf []byte) *OffsetByteBuffer {
	return &OffsetByteBuffer{
		Buffer: bytes.Buffer{},
	}
}

// 这里没有存成文件，最好基于文件进行传输
func Handler(ctx context.Context, event events.S3Event) (result string, err error) {
	for _, eventRecord := range event.Records {
		bucket := eventRecord.S3.Bucket
		fileName := eventRecord.S3.Object.Key
		byteBuffer := NewOffsetByteBuffer(nil)
		err = DownloadOriginalImage(byteBuffer, bucket.Name, fileName)
		if nil != err {
			logrus.Errorf("download original image failed. %s", err)
			return
		}

		resizedWriter := bytes.NewBuffer(nil)
		err = ResizeImage(byteBuffer, resizedWriter, 100)
		if nil != err {
			logrus.Errorf("resize image failed. %s", err)
			return
		}

		err = UploadThumbnailImage(resizedWriter, "xl-pic-resized", fmt.Sprintf("%d-%s", 100, fileName))
		if nil != err {
			logrus.Errorf("upload thumbnail image failed. %s", err)
			return
		}
	}

	result = "success"
	return
}

func main() {
	lambda.Start(Handler)
}
