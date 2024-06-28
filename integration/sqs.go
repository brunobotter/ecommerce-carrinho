package integration

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/external"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/brunobotter/ecommerce-carrinho/configs"
)

func SendMessageToSQS(queueURL string, messageBody string) error {
	logger := configs.GetLogger("SQS")

	// Configurar as credenciais na configuração AWS SDK v2
	cfg, err := external.LoadDefaultAWSConfig()
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
		MessageGroupId: aws.String("paymentGroup"), // Fila FIFO requer um ID de grupo
	}

	// Envia a mensagem
	result, err := svc.SendMessage(context.TODO(), sendMessageInput)
	if err != nil {
		logger.Errorf("failed to send message to SQS: %v", err)
		return err
	}

	logger.Infof("Message sent, ID: %s", *result.MessageId)
	return nil
}

func GetAWSSecrets() (aws.CredentialsProvider, error) {
	// Inicializar cliente Secrets Manager
	cfg, err := external.LoadDefaultAWSConfig()
	if err != nil {
		return nil, err
	}

	svc := secretsmanager.NewFromConfig(cfg)

	secretName := "myapp" // Nome do seu segredo no AWS Secrets Manager
	versionStage := "AWSCURRENT"

	// Preparar input para obter o valor do segredo
	input := &secretsmanager.GetSecretValueInput{
		SecretId:     aws.String(secretName),
		VersionStage: aws.String(versionStage),
	}

	// Obter o valor do segredo
	result, err := svc.GetSecretValue(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	// Decodificar o valor do segredo
	var creds aws.CredentialsProvider
	if err := json.Unmarshal([]byte(*result.SecretString), &creds); err != nil {
		return nil, err
	}

	return creds, nil
}
