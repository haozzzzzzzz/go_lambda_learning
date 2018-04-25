package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"ExampleApi/constant"
	"ExampleApi/handler"

	"github.com/haozzzzzzzz/go-lambda/resource/iam"
	"github.com/sirupsen/logrus"
)

func main() {
	var err error
	filePath := flag.String("path", "", "role.yaml save file path")
	flag.Parse()

	if *filePath == "" {
		logrus.Errorf("wrong role save target file path")
		return
	}

	handler.GetMainHandler()

	roleFilePath, err := filepath.Abs(*filePath)
	if nil != err {
		logrus.Errorf("get absolute file path failed. \n%s.", err)
		return
	}

	logrus.Info("detecting project")
	defer func() {
		logrus.Info("detecting project finish")
	}()

	role := iam.GetExecutionRole(fmt.Sprintf("%sRole", constant.LambdaFunctionName))
	err = role.WriteTo(roleFilePath, os.ModePerm)
	if nil != err {
		logrus.Errorf("write role to yaml failed. \n%s.", err)
		return
	}
}
