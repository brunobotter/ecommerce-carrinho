package configs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var awsConfig *aws.Config

func GetConfig() *aws.Config {
	if awsConfig == nil {
		cfg, err := config.LoadDefaultConfig(context.TODO())
		if err != nil {
			panic("unable to connect to AWS")
		}

		awsConfig = &cfg
	}
	return awsConfig
}
