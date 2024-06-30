package integration

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/brunobotter/ecommerce-carrinho/configs"
)

func SendMessageToSQS(queueURL string, messageBody string) error {
	logger := configs.GetLogger("SQS")
	ctx := context.Background()

	// Carregar configuração padrão do SDK AWS com base em variáveis de ambiente
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion("us-east-1"))
	if err != nil {
		logger.Errorf("failed to load SDK configuration: %v", err)
		return err
	}

	sqsClient := sqs.NewFromConfig(cfg)

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
