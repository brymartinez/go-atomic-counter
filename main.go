package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func main() {
	var ctx context.Context
	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion("ap-southeast-1"),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...interface{}) (aws.Endpoint, error) {
				return aws.Endpoint{URL: "http://localhost:4566"}, nil
			}),
		),
	)

	client := dynamodb.NewFromConfig(cfg)

	updateField := "Refrigerator"
	// Define the parameters for the update operation
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String("IDGenerator-local"),
		Key: map[string]types.AttributeValue{
			"YourPartitionKey": &types.AttributeValueMemberS{
				Value: "IDSequence",
			},
		},
		UpdateExpression: aws.String("SET" + updateField + " = " + updateField + ":value"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":value": &types.AttributeValueMemberN{
				Value: "1",
			},
		},
		ReturnValues: types.ReturnValueAllNew,
	}

	_, err = client.UpdateItem(context.Background(), input)
	if err != nil {
		log.Fatalf("Got error calling UpdateItem: %s", err)
	}

}
