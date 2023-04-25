package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

func main() {
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
	data, err := sess.Config.Credentials.Get()
	fmt.Println(data)
}

//ASIARBMYK7AQ4N4OO3O2
//JlS7o2Snf+mubTbRCD3SB9vK/eLAZpwKd1u7e9H4
