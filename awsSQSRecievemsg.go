package main

//Necessary packages are imported here.
import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// GetQueueURL function.
func GetQueueURL(sess *session.Session, queue string) (*sqs.GetQueueUrlOutput, error) {
	sqsClient := sqs.New(sess)
	//To get the URL of the specified queue.
	result, err := sqsClient.GetQueueUrl(&sqs.GetQueueUrlInput{
		QueueName: &queue,
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

// ReceiveMessage Function.
func GetMessages(sess *session.Session, queueUrl string, maxMessages int) (*sqs.ReceiveMessageOutput, error) {
	sqsClient := sqs.New(sess)

	msgResult, err := sqsClient.ReceiveMessage(&sqs.ReceiveMessageInput{
		QueueUrl:            &queueUrl,
		MaxNumberOfMessages: aws.Int64(1),
	})

	if err != nil {
		return nil, err
	}

	return msgResult, nil
}

func main() {
	//Create a new AWS session with some options like AWS profile and region.
	session, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region: aws.String("us-east-1"),
		},
	})

	if err != nil {
		fmt.Printf("Failed to initialize new session: %v", err)
		return
	}
	//Declared necessary variable to specify the Queuename.
	queueName := "Bhuvana-firstqueue"
	//URL function is called.
	urlRes, err := GetQueueURL(session, queueName)
	if err != nil {
		fmt.Printf("Got an error while trying to create queue: %v", err)
		return
	}
	// To print the url
	fmt.Println("Got Queue URL: " + *urlRes.QueueUrl)

	maxMessages := 1
	msgRes, err := GetMessages(session, *urlRes.QueueUrl, maxMessages)
	if err != nil {
		fmt.Printf("Got an error while trying to retrieve message: %v", err)
		return
	}
	//It will pritn the message body and ReceiptHandle(used to delete and modify the msg later).
	fmt.Println("Message Body: " + *msgRes.Messages[0].Body)
	fmt.Println("Message Handle: " + *msgRes.Messages[0].ReceiptHandle)

}
