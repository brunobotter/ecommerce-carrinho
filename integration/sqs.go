package integration

import (
	"crypto/tls"
	"net/http"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/brunobotter/ecommerce-carrinho/configs"
)

func SendMessageToSQS(queueURL string, messageBody string, accessKey string, secretAccess string, region string) error {
	logger := configs.GetLogger("SQS")

	// Obter as credenciais do GitHub Secrets
	accessKeyID := accessKey
	secretAccessKey := secretAccess

	// Criar uma sessão AWS usando as credenciais obtidas
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(region),
		Credentials: credentials.NewStaticCredentials(
			accessKeyID,
			secretAccessKey,
			"", // token (optional)
		),
	})
	if err != nil {
		return err
	}
	awsHTTPClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true}, // Ignorar a verificação de certificado
		},
		Timeout: 10 * time.Second, // Ajuste o timeout conforme necessário
	}

	sess.Config.HTTPClient = awsHTTPClient
	sqsClient := sqs.New(sess)

	// Parâmetros da mensagem
	sendMessageInput := &sqs.SendMessageInput{
		MessageBody:    aws.String(messageBody),
		QueueUrl:       aws.String(queueURL),
		MessageGroupId: aws.String("paymentGroup"),
	}

	// Enviar mensagem
	result, err := sqsClient.SendMessage(sendMessageInput)
	if err != nil {
		logger.Errorf("failed to send message to SQS: %v", err)
		return err
	}

	logger.Infof("Message sent, ID: %s", *result.MessageId)
	return nil
}
