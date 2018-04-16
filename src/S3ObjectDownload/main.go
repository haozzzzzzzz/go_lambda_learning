package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	bucket := "xl-pic"
	item := "HappyFace.jpg"

	file, err := os.Create(fmt.Sprintf("/Users/hao/Documents/Projects/Github/go_lambda_learning/src/S3PicResizeHandler/%s", item))
	if nil != err {
		exitError("unable to open file %q, %v", err)
		return
	}
	defer file.Close()

	ses, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	downloader := s3manager.NewDownloader(ses)

	numBytes, err := downloader.Download(file, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(item),
	})

	if nil != err {
		exitError("Unable to download item %q, %v", item, err)
		return
	}

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes")
}

func exitError(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
