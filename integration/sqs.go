package integration

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/brunobotter/ecommerce-carrinho/config"
)

func SendMessageToSQS(queueURL string, messageBody string) error {
	logger := config.GetLogger("SQS")
	// Cria uma sessão AWS

	sess, err := session.NewSession(&aws.Config{
		Region:                        aws.String("us-east-1"),
		CredentialsChainVerboseErrors: aws.Bool(true),
	})
	if err != nil {
		logger.Errorf("failed to create session: %w", err)
		return err
	}

	// Cria um serviço SQS
	svc := sqs.New(sess)

	// Parâmetros da mensagem
	sendMessageInput := &sqs.SendMessageInput{
		MessageBody:    aws.String(messageBody),
		QueueUrl:       aws.String(queueURL),
		MessageGroupId: aws.String("paymentGroup"), // Fila FIFO requer um ID de grupo
	}

	// Envia a mensagem
	result, err := svc.SendMessage(sendMessageInput)
	if err != nil {
		logger.Errorf("failed to send message to SQS: %w", err)
		return err
	}

	fmt.Println("Message ID", *result.MessageId)
	return nil
}
