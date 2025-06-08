package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
	"github.com/google/uuid"
)

var awsConfig *aws.Config

func main() {
	ctx := context.Background()

	sqsClient := sqs.NewFromConfig(*getConfig())

	result, err := sqsClient.GetQueueUrl(ctx, &sqs.GetQueueUrlInput{
		QueueName: aws.String("test-sqs-queue"), // name of the sqs queue
	})
	if err != nil {
		fmt.Println("Unable to connect to SQS: ", err)
	}

	fmt.Println("Queue URL: ", *result.QueueUrl)

	// SendMessage(ctx, *sqsClient, result.QueueUrl)
	// SendBulkMessage(ctx, *sqsClient, result.QueueUrl)
	ReceiveMessage(ctx, *sqsClient, result.QueueUrl)
	// ReceiveMessageBatch(ctx, *sqsClient, result.QueueUrl)
	// DeleteMessage(ctx, *sqsClient, result.QueueUrl, types.Message{})
	// DeleteMessageBatch(ctx, *sqsClient, result.QueueUrl, []types.Message{})
	// PurgeMessages(ctx, *sqsClient, result.QueueUrl)

}

func SendMessage(ctx context.Context, sqsClient sqs.Client, queueUrl *string) {
	_, errr := sqsClient.SendMessage(ctx, &sqs.SendMessageInput{
		QueueUrl:    queueUrl,
		MessageBody: aws.String("Sending Message to SQS from code"),
	})
	if errr != nil {
		fmt.Println("Unable to send message to SQS: ", errr)
	} else {
		fmt.Println("Message successfully sent to SQS")
	}
}

func SendBulkMessage(ctx context.Context, sqsClient sqs.Client, queueUrl *string) {
	messageBatch := make([]types.SendMessageBatchRequestEntry, 10)
	for i := 0; i < 10; i++ {
		message := fmt.Sprintf("Batch Message Number %d", i)
		messageBatch[i] = types.SendMessageBatchRequestEntry{
			Id:          aws.String(uuid.New().String()),
			MessageBody: &message,
		}
	}

	_, errr := sqsClient.SendMessageBatch(ctx, &sqs.SendMessageBatchInput{
		QueueUrl: queueUrl,
		Entries:  messageBatch,
	})
	if errr != nil {
		fmt.Println("Unable to send batch messages to SQS: ", errr)
	} else {
		fmt.Println("Batch Messages sent successfully to SQS")
	}
}

func ReceiveMessage(ctx context.Context, sqsClient sqs.Client, queueUrl *string) {
	// This will constantly listen to the SQS queue and print the messages
	for {
		messagesResult, err := sqsClient.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
			QueueUrl: queueUrl,
		})
		if err != nil {
			fmt.Println("Unable to receive mesaages from SQS: ", err)
		}

		for _, message := range messagesResult.Messages {
			fmt.Println("Message: ", *message.Body)
			DeleteMessage(ctx, sqsClient, queueUrl, message)
		}
	}
}

func ReceiveMessageBatch(ctx context.Context, sqsClient sqs.Client, queueUrl *string) {
	var messageBatch []types.Message
	messagesResult, err := sqsClient.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:            queueUrl,
		VisibilityTimeout:   5,
		MaxNumberOfMessages: 5,
	})
	if err != nil {
		fmt.Println("Unable to receive mesaages from SQS: ", err)
	}
	for _, message := range messagesResult.Messages {
		fmt.Println("Message Body: ", *message.Body)
		messageBatch = append(messageBatch, message)
	}
	DeleteMessageBatch(ctx, sqsClient, queueUrl, messageBatch)
}

func DeleteMessage(ctx context.Context, sqsClient sqs.Client, queueUrl *string, message types.Message) {
	_, errr := sqsClient.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      queueUrl,
		ReceiptHandle: message.ReceiptHandle,
	})
	if errr != nil {
		fmt.Println("Unable to delete message from SQS: ", errr)
	} else {
		fmt.Println("Message successfully deleted from SQS")
	}
}

func DeleteMessageBatch(ctx context.Context, sqsClient sqs.Client, queueUrl *string, messages []types.Message) {
	var messageBatch []types.DeleteMessageBatchRequestEntry
	for _, message := range messages {
		messageBatch = append(messageBatch, types.DeleteMessageBatchRequestEntry{
			Id:            message.MessageId,
			ReceiptHandle: message.ReceiptHandle,
		})
	}
	_, errr := sqsClient.DeleteMessageBatch(ctx, &sqs.DeleteMessageBatchInput{
		QueueUrl: queueUrl,
		Entries:  messageBatch,
	})
	if errr != nil {
		fmt.Println("Unable to delete messages in batch from SQS: ", errr)
	} else {
		fmt.Println("Message successfully deleted in batch from SQS")
	}
}

func PurgeMessages(ctx context.Context, sqsClient sqs.Client, queueUrl *string) {
	_, errr := sqsClient.PurgeQueue(ctx, &sqs.PurgeQueueInput{
		QueueUrl: queueUrl,
	})
	if errr != nil {
		fmt.Println("Unable to purge messages from SQS: ", errr)
	} else {
		fmt.Println("Messages successfully purged from SQS")
	}
}

func getConfig() *aws.Config {
	if awsConfig == nil {
		cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion("eu-north-1"))
		if err != nil {
			panic("Unable to connect to AWS")
		}
		awsConfig = &cfg
		return awsConfig
	}
	return awsConfig
}
