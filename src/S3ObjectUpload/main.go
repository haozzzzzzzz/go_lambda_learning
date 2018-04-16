package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	bucket := "xl-pic-resized"
	filename := "./src/S3ObjectUpload/HappyFace.jpg"

	file, err := os.Open(filename)
	if nil != err {
		exitError("Unable to open file %q, %v", err)
		return
	}
	defer file.Close()

	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})

	uploader := s3manager.NewUploader(sess)
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String("new_HappyFace.jpg"),
		Body:   file,
	})
	if nil != err {
		exitError("Unable to upload %q to %q\n", filename, bucket)
		return
	}

	fmt.Printf("Successfully uploaded %q to %q\n", filename, bucket)
}

func exitError(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
