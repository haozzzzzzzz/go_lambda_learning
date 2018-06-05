package config

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/go-playground/validator"
	"github.com/haozzzzzzzz/go-rapid-development/utils/yaml"
	"github.com/sirupsen/logrus"
)

type AWSConfigFormat struct {
	Region string `yaml:"region" validate:"required"`
}

var AWSConfig AWSConfigFormat
var AWSSession *session.Session

func init() {
	var err error
	err = yaml.ReadYamlFromFile("./config/aws.yaml", &AWSConfig)
	if nil != err {
		logrus.Fatalf("read aws config file failed. %s", err)
		return
	}

	err = validator.New().Struct(AWSConfig)
	if nil != err {
		logrus.Fatalf("validate aws config file failed. %s", err)
		return
	}

	awsConfig := &aws.Config{}
	awsConfig.Region = aws.String(AWSConfig.Region)
	AWSSession, err = session.NewSession(awsConfig)
	if nil != err {
		logrus.Fatalf("new aws session failed. %s", err)
		return
	}

}
