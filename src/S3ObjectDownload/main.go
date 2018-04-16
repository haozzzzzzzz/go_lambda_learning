package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/sirupsen/logrus"
)

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

func main() {
	bucket := "xl-pic"
	item := "HappyFace.jpg"

	// 相对项目根目录
	file, err := os.Create(fmt.Sprintf("./src/S3ObjectDownload/%s", item))
	if nil != err {
		exitError("unable to open file %q, %v", err)
		return
	}
	defer file.Close()

	ses, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	downloader := s3manager.NewDownloader(ses)

	byteBuffer := NewOffsetByteBuffer(nil)
	numBytes, err := downloader.Download(byteBuffer, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(item),
	})
	if nil != err {
		exitError("Unable to download item %q, %v", item, err)
		return
	}

	_, err = byteBuffer.WriteTo(file)
	if nil != err {
		exitError("Write data to file failed. %s", err)
		return
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
}

func exitError(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
