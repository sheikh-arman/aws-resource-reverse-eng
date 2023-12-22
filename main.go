package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"

	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	// Create a new session using your credentials and region
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	// Create DynamoDB client
	svc := dynamodb.New(sess)

	// List all DynamoDB tables
	input := &dynamodb.ListTablesInput{}
	result, err := svc.ListTables(input)
	if err != nil {
		fmt.Println("Error listing tables:", err)
		return
	}

	fmt.Println("Tables:")
	for _, tableName1 := range result.TableNames {
		fmt.Println(*tableName1)
		tableName := *tableName1
		input := &dynamodb.DescribeTableInput{
			TableName: aws.String(tableName),
		}
		result, err := svc.DescribeTable(input)
		if err != nil {
			fmt.Println("Error describing table:", err)
			return
		}

		// Print details of the table
		fmt.Println("Table Description:")
		fmt.Println("Table Name:", *result.Table.TableName)
		fmt.Println("Attribute Definitions:")
		for _, attr := range result.Table.AttributeDefinitions {
			fmt.Println("  Name:", *attr.AttributeName)
			fmt.Println("  Type:", *attr.AttributeType)
		}
		fmt.Println("Key Schema:")
		for _, key := range result.Table.KeySchema {
			fmt.Println("  Name:", *key.AttributeName)
			fmt.Println("  Type:", *key.KeyType)
		}
		fmt.Println(result)
	}
}
