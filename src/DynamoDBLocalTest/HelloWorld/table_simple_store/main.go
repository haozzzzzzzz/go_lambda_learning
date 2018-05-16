package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/haozzzzzzzz/go-lambda/resource/dyndb/table"
)

func main() {
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("ap-south-1"),
		Endpoint: aws.String("http://0.0.0.0:8000"),
	})
	if nil != err {
		log.Fatal(err)
		return
	}
	svc := dynamodb.New(sess)
	simpleStore := table.SimpleStoreTable{
		DynamoDBTable: table.DynamoDBTable{
			TableName: "video_buddy_simple_store_dev",
			Client:    svc,
			Ctx:       aws.BackgroundContext(),
		},
	}
	obj := &table.SimpleStoreModel{
		PartitionKey: "video_buddy_cracked_tags",
		SortKey:      "1",
		Value: struct {
			Name    string `json:"name"`
			Address string `json:"address"`
		}{
			Name:    "Haozi",
			Address: "Shenzhen",
		},
	}
	err = simpleStore.AddSimpleStore(obj)
	if nil != err {
		log.Fatal(err)
		return
	}

	newObj, err := simpleStore.GetSimpleStore(obj.PartitionKey, obj.SortKey)
	if nil != err {
		log.Fatal(err)
		return
	}

	fmt.Println(newObj)

	err = simpleStore.DeleteSimpleStore(obj.PartitionKey, obj.SortKey)
	if nil != err {
		log.Fatal(err)
		return
	}
}
