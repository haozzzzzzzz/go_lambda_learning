package handler

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	kinesis2 "github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/aws/aws-xray-sdk-go/xray"
	"github.com/haozzzzzzzz/go-lambda/resource/kinesis"
	"github.com/haozzzzzzzz/go-rapid-development/utils/id"
	"github.com/sirupsen/logrus"
)

func BasicExecutionEventHandler(ctx context.Context, event interface{}) (msg string, err error) {
	xray.Configure(xray.Config{LogLevel: "trace"})

	defer func() {
		if nil != err {
			msg = err.Error()
			return
		}
	}()

	svc, err := kinesis.NewSimpleKinesis("ap-south-1")
	if nil != err {
		logrus.Errorf("new simple kinesis failed. %s.", err)
		return
	}

	output, err := svc.PutRecord(&kinesis2.PutRecordInput{
		Data:         []byte("Hello"),
		StreamName:   aws.String("KinesisTest"),
		PartitionKey: aws.String(id.UniqueID()),
	})
	if nil != err {
		logrus.Errorf("put record failed. %s.", err)
		return
	}

	msg = fmt.Sprintf("%#v", output)
	return
}
