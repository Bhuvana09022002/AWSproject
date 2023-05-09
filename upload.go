package main

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"fmt"
)

// Uploadfile function.
func UploadFile(uploader *s3manager.Uploader, filePath string, bucketName string, fileName string) error {
	//To open the file from local system using filepath.
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}

	defer file.Close()
	//Uploader struct used to upload the files into the bucket.
	_, err = uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(fileName),
		Body:   file,
	})

	return err
}

// The main function starts here.
func main() {
	//Create a new AWS session with some options like AWS profile and region.
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "default",
		Config: aws.Config{
			Region: aws.String("us-east-1"),
		},
	})

	if err != nil {
		fmt.Printf("Failed to initialize new session: %v", err)
		return
	}
	//Declared necessary variables.
	bucketName := "bhuvana-firstbucket"
	uploader := s3manager.NewUploader(sess)
	filename := "flower.txt"

	//Uploadfile function is called.
	err = UploadFile(uploader, "flower.txt", bucketName, filename)
	if err != nil {
		fmt.Printf("Failed to upload file: %v", err)
	}
	//If the file is sucessfully uploaded.
	fmt.Println("Successfully uploaded file!")
}
