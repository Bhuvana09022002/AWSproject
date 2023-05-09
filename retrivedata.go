package main

//Necessary packages are imported here.
import (
	"fmt"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

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

	// Create an S3 client
	svc := s3.New(session)

	// //Declared necessary variables to specify the bucket and object key.
	bucket := "bhuvana-firstbucket"
	key := "flower.txt"

	// Retrieve the object data
	object, err := svc.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		fmt.Println("Error retrieving object data:", err)
		return
	}

	// Read the object data into a byte slice
	data := make([]byte, 0, *object.ContentLength)
	buf := make([]byte, 1024)
	for {
		n, err := object.Body.Read(buf)
		if err != nil && err != io.EOF {
			fmt.Println("Error reading object data:", err)
			return
		}
		if n == 0 {
			break
		}
		data = append(data, buf[:n]...)
	}

	// Print the object data
	fmt.Println(string(data))
}
