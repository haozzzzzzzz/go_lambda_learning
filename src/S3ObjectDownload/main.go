package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func main() {
	if len(os.Args) != 3 {
		exitError("Bucket and item names required\nUsage: %s bucket_name item_name", os.Args[0])
	}

	bucket := os.Args[1]
	item := os.Args[2]

	file, err := os.Create(item)
	if nil != err {
		exitError("unable to open file %q, %v", err)
		return
	}
	defer file.Close()

	fmt.Print(bucket)

	ses, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1"),
	})
	downloader := s3manager.NewDownloader(ses)

}

func exitError(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
