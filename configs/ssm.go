package configs

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
)

var (
	ssmClient *secretsmanager.Client
)

func InitSSM(region string) error {
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(region))
	if err != nil {
		return err
	}

	ssmClient = secretsmanager.NewFromConfig(cfg)
	return nil
}

func GetSSMClient() *secretsmanager.Client {
	return ssmClient
}
