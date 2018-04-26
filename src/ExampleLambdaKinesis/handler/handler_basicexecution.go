package handler

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	kinesis2 "github.com/aws/aws-sdk-go/service/kinesis"
	"github.com/haozzzzzzzz/go-lambda/resource/kinesis"
	"github.com/sirupsen/logrus"
)

func BasicExecutionEventHandler(ctx context.Context, event interface{}) (msg string, err error) {
	go func() {
		svc, err := kinesis.GetKinesis("ap-south-1")
		if nil != err {
			msg = err.Error()
			logrus.Errorf("get kinesis failed. %s", err)
			return
		}

		var streamName string = "KinesisTest"
		putOutput, err := svc.PutRecord(&kinesis2.PutRecordInput{
			Data:         []byte("hello"),
			StreamName:   &streamName,
			PartitionKey: aws.String("55555"),
		})
		if nil != err {
			logrus.Error("put record failed. %s", err)
			return
		}

		msg = fmt.Sprintf("%#v", putOutput)
	}()
	msg = "成功调用"
	return
}
