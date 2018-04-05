package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Testing Lambda
// func hello() (string, error) {
// 	return "Hello ƛ!", nil
// }

// func HandleRequest(ctx context.Context, name MyEvent) (string, error) {
// 	return fmt.Sprintf("Hello %s!", name.Name), nil
// }

// --------------------------------------------------------------
// Testing Lambda function

type MyEvent struct {
	// Name string `json:"name"`
	Name string `json:"What is your name?"`
	Age  int    `json:"How old are you?"`
}

type MyResponse struct {
	Message string `json:"Answer:"`
	Ok      bool   `json:"ok"`
}

func HandleLambdaEvent(event MyEvent) (MyResponse, error) {
	return MyResponse{
		Message: fmt.Sprintf("%s is %d years old!", event.Name, event.Age),
		Ok:      true,
	}, nil
}

// --------------------------------------------------------------
// Testing S3 function

var invokeCount = 0
var myObjects []*s3.Object

func init() {
	svc := s3.New(session.New())
	input := &s3.ListObjectsV2Input{
		Bucket: aws.String("examplebucket"),
	}
	result, _ := svc.ListObjectsV2(input)
	myObjects = result.Contents
}

func LambdaHandler() int {
	invokeCount = invokeCount + 1
	log.Print(myObjects)
	return invokeCount
}

func main() {
	lambda.Start(LambdaHandler)
}
