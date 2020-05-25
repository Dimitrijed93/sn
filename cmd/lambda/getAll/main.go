package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Response typedef
type Response events.APIGatewayProxyResponse

type snippet struct {
	id      int
	title   string
	content string
	created string
}

// Handler is our lambda handler invoked by the `lambda.Start` function call
func Handler(ctx context.Context) (Response, error) {

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	svc := dynamodb.New(sess)

	tableName := "snippets"

	params := &dynamodb.ScanInput{
		TableName: aws.String(tableName),
	}

	var snippets []snippet

	result, err := svc.Scan(params)

	if err != nil {
		fmt.Println("Got error while scanning:")
		fmt.Println(err.Error())
		os.Exit(1)
	}

	for _, i := range result.Items {
		snippet := snippet{}
		err := dynamodbattribute.UnmarshalMap(i, &snippet)
		if err != nil {
			fmt.Println("Got error unmarshalling: dynamodbattribute")
			fmt.Println(err.Error())
			os.Exit(1)
		}
		snippets = append(snippets, snippet)
	}

	var buf bytes.Buffer

	body, err := json.Marshal(snippets)
	if err != nil {
		fmt.Println("Got error while marshaling response:")
		fmt.Println(err.Error())
		os.Exit(1)
	}
	json.HTMLEscape(&buf, body)

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}

	return resp, nil
}

func main() {
	lambda.Start(Handler)
}
