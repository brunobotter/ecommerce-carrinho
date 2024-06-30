package integration

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/brunobotter/ecommerce-carrinho/configs"
)

func SendMessageToSQS(queueURL string, messageBody string, accessKey string, secretAccess string, region string) error {
	logger := configs.GetLogger("SQS")
	ctx := context.Background()

	// Obter as credenciais do GitHub Secrets
	accessKeyID := accessKey
	secretAccessKey := secretAccess

	// Criar uma sessão AWS usando as credenciais obtidas
	sess, err := session.NewSession(&aws.Config{
		Region: region,
		Credentials: credentials.NewStaticCredentials(
			accessKeyID,
			secretAccessKey,
			"", // token (optional)
		),
	})
	if err != nil {
		return err
	}

	sqsClient := sqs.NewFromConfig(sess)

	logger.Debugf("iniciou service sqs")
	// Parâmetros da mensagem
	sendMessageInput := &sqs.SendMessageInput{
		MessageBody:    aws.String(messageBody),
		QueueUrl:       aws.String(queueURL),
		MessageGroupId: aws.String("paymentGroup"),
	}
	logger.Debugf("enviando msg")
	// Enviar mensagem
	result, err := sqsClient.SendMessage(ctx, sendMessageInput)
	if err != nil {
		logger.Errorf("failed to send message to SQS: %v", err)
		return err
	}

	logger.Infof("Message sent, ID: %s", *result.MessageId)
	return nil
}
