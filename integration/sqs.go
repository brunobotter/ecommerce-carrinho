package integration

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/brunobotter/ecommerce-carrinho/configs"
)

func SendMessageToSQS(queueURL string, messageBody string) error {
	logger := configs.GetLogger("SQS")

	// Obter as credenciais do AWS Secrets Manager
	creds, err := GetAWSSecrets()
	if err != nil {
		logger.Errorf("failed to retrieve AWS credentials from Secrets Manager: %v", err)
		return err
	}

	// Carregar configurações padrão do AWS SDK v2 com credenciais
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(creds),
	)
	if err != nil {
		logger.Errorf("failed to load AWS SDK configuration: %v", err)
		return err
	}

	// Criar um cliente SQS
	svc := sqs.NewFromConfig(cfg)

	// Parâmetros da mensagem
	sendMessageInput := &sqs.SendMessageInput{
		MessageBody:    aws.String(messageBody),
		QueueUrl:       aws.String(queueURL),
		MessageGroupId: aws.String("paymentGroup"),
	}

	// Enviar mensagem
	result, err := svc.SendMessage(context.TODO(), sendMessageInput)
	if err != nil {
		logger.Errorf("failed to send message to SQS: %v", err)
		return err
	}

	logger.Infof("Message sent, ID: %s", *result.MessageId)
	return nil
}

func GetAWSSecrets() (aws.CredentialsProvider, error) {
	// Carregar configurações padrão do AWS SDK v2
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return nil, err
	}

	// Criar cliente Secrets Manager
	svc := secretsmanager.NewFromConfig(cfg)

	secretName := "myapp" // Nome do seu segredo no AWS Secrets Manager
	versionStage := "AWSCURRENT"

	// Preparar input para obter o valor do segredo
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String(versionStage),
	}

	// Obter valor do segredo
	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	// Estrutura para armazenar as credenciais
	type Secret struct {
		AccessKeyID     string `json:"accessKeyId"`
		SecretAccessKey string `json:"secretAccessKey"`
		SessionToken    string `json:"sessionToken"`
	}

	var secret Secret
	if err := json.Unmarshal([]byte(*result.SecretString), &secret); err != nil {
		return nil, err
	}

	// Criar um provedor de credenciais usando as credenciais obtidas
	credsProvider := aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
		return aws.Credentials{
			AccessKeyID:     secret.AccessKeyID,
			SecretAccessKey: secret.SecretAccessKey,
			SessionToken:    secret.SessionToken,
		}, nil
	})

	return credsProvider, nil
}
