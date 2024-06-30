package integration

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/brunobotter/ecommerce-carrinho/configs"
)

func SendMessageToSQS(queueURL string, messageBody string) error {
	logger := configs.GetLogger("SQS")

	// Cria uma sessão AWS
	sess := session.Must(session.NewSession(&aws.Config{
		Region:   aws.String("us-east-1"),
		Endpoint: aws.String(""),
	}))

	// Criar um cliente SQS
	svc := sqs.New(sess)

	// Parâmetros da mensagem
	sendMessageInput := &sqs.SendMessageInput{
		MessageBody:    aws.String(messageBody),
		QueueUrl:       aws.String(queueURL),
		MessageGroupId: aws.String("paymentGroup"),
	}

	// Enviar mensagem
	result, err := svc.SendMessage(sendMessageInput)
	if err != nil {
		logger.Errorf("failed to send message to SQS: %v", err)
		return err
	}

	logger.Infof("Message sent, ID: %s", *result.MessageId)
	return nil
}
