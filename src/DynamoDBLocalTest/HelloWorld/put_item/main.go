package main

import (
	"fmt"
	"log"
	"time"

	"DynamoDBLocalTest/HelloWorld/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	dynamodb2 "github.com/haozzzzzzzz/go-lambda/resource/dynamodb"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("ap-south-1"),
		Endpoint: aws.String("http://192.168.0.36:8000"),
	})
	if nil != err {
		log.Fatal(err)
		return
	}
	svc := dynamodb.New(sess)
	client := dynamodb2.DynamoDBTable{
		TableName: "video_buddy_home_dev",
		Client:    svc,
	}
	_ = client

	now := time.Now()

	homes := []model.Home{
		{
			HomeId:              1,
			Title:               "首页1",
			Style:               0,
			AppVersionMin:       "1.0",
			AppVersionMax:       "2.0",
			EffectiveTime:       now,
			OnlineState:         model.HomeOnlineStateDefault,
			LastUpdateUid:       "1",
			LastUpdateUsername:  "罗浩",
			LastUpdateTime:      now,
			LastUpdateTimestamp: now.Unix(),
			DeviceIds:           []string{"0", "a"},
			CreateTime:          now,
		},
		{
			HomeId:              2,
			Title:               "首页1",
			Style:               0,
			AppVersionMin:       "1.0",
			AppVersionMax:       "3.0",
			EffectiveTime:       now,
			OnlineState:         model.HomeOnlineStateDefault,
			LastUpdateUid:       "1",
			LastUpdateUsername:  "罗浩",
			LastUpdateTime:      now,
			LastUpdateTimestamp: now.Unix(),
			DeviceIds:           []string{"c", "1"},
			CreateTime:          now,
		},
		{
			HomeId:              3,
			Title:               "首页1",
			Style:               0,
			AppVersionMin:       "3.0",
			AppVersionMax:       "",
			EffectiveTime:       now,
			OnlineState:         model.HomeOnlineStateDefault,
			LastUpdateUid:       "1",
			LastUpdateUsername:  "罗浩",
			LastUpdateTime:      now,
			LastUpdateTimestamp: now.Unix(),
			DeviceIds:           []string{},
			CreateTime:          now,
		},
	}

	for _, home := range homes {
		fmt.Println(home)
		err := client.PutItem(home)
		if nil != err {
			log.Fatal(err)
			return
		}
	}

}
