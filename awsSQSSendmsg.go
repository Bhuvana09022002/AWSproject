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

// SendMessage Function.
func SendMessage(sess *session.Session, queueUrl string, messageBody string) error {
	sqsClient := sqs.New(sess)

	_, err := sqsClient.SendMessage(&sqs.SendMessageInput{
		QueueUrl:    &queueUrl,
		MessageBody: aws.String(messageBody),
	})

	return err
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

	// The message for the queue is given here.
	messageBody := "This is a test message"
	//Message funtion is called here.
	err = SendMessage(session, *urlRes.QueueUrl, messageBody)
	if err != nil {
		fmt.Printf("Got an error while trying to send message to queue: %v", err)
		return
	}

	fmt.Println("Message sent successfully")
}
