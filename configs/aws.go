package configs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var awsConfig *aws.Config

type Connection struct {
	config *aws.Config
}

func NewConnection() *Connection {
	return &Connection{config: GetConfig()}
}

func GetConfig() *aws.Config {
	if awsConfig == nil {
		awsConfig, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			panic("unable to connect to AWS")
		}

		return &awsConfig
	}
	return awsConfig
}
