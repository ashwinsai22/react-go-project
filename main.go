package main

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
)

type Todo struct {
	ID        string `json:"id" dynamodbav:"id"`
	Completed bool   `json:"completed" dynamodbav:"completed"`
	Body      string `json:"body" dynamodbav:"body"`
}

var db *dynamodb.Client
var tableName = "todos"

func init() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal("Unable to load AWS config:", err)
	}
	db = dynamodb.NewFromConfig(cfg)
}

func handler(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	switch request.RouteKey {

	case "GET /api/todos":
		return getTodos(ctx)

	case "POST /api/todos":
		return createTodo(ctx, request)

	case "PATCH /api/todos/{id}":
		id := request.PathParameters["id"]
		return updateTodo(ctx, id)

	case "DELETE /api/todos/{id}":
		id := request.PathParameters["id"]
		return deleteTodo(ctx, id)

	case "OPTIONS /api/todos":
		return response(200, "")

	case "OPTIONS /api/todos/{id}":
		return response(200, "")

	default:
		return response(404, `{"error":"Not Found"}`)
	}
}

func getTodos(ctx context.Context) (events.APIGatewayV2HTTPResponse, error) {

	result, err := db.Scan(ctx, &dynamodb.ScanInput{
		TableName: &tableName,
	})
	if err != nil {
		return response(500, err.Error())
	}

	var todos []Todo
	err = attributevalue.UnmarshalListOfMaps(result.Items, &todos)
	if err != nil {
		return response(500, err.Error())
	}

	body, _ := json.Marshal(todos)
	return response(200, string(body))
}

func createTodo(ctx context.Context, request events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {

	var todo Todo
	if err := json.Unmarshal([]byte(request.Body), &todo); err != nil {
		return response(400, "Invalid request body")
	}

	if todo.Body == "" {
		return response(400, "Todo body cannot be empty")
	}

	todo.ID = uuid.New().String()
	todo.Completed = false

	item, err := attributevalue.MarshalMap(todo)
	if err != nil {
		return response(500, err.Error())
	}

	_, err = db.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: &tableName,
		Item:      item,
	})
	if err != nil {
		return response(500, err.Error())
	}

	body, _ := json.Marshal(todo)
	return response(201, string(body))
}

func updateTodo(ctx context.Context, id string) (events.APIGatewayV2HTTPResponse, error) {

	_, err := db.UpdateItem(ctx, &dynamodb.UpdateItemInput{
		TableName: &tableName,
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
		UpdateExpression: awsString("SET completed = :val"),
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":val": &types.AttributeValueMemberBOOL{Value: true},
		},
	})
	if err != nil {
		return response(500, err.Error())
	}

	return response(200, `{"success":true}`)
}

func deleteTodo(ctx context.Context, id string) (events.APIGatewayV2HTTPResponse, error) {

	_, err := db.DeleteItem(ctx, &dynamodb.DeleteItemInput{
		TableName: &tableName,
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		return response(500, err.Error())
	}

	return response(200, `{"success":true}`)
}

func response(status int, body string) (events.APIGatewayV2HTTPResponse, error) {
	return events.APIGatewayV2HTTPResponse{
		StatusCode: status,
		Headers: map[string]string{
			"Content-Type":                 "application/json",
			"Access-Control-Allow-Origin":  "*",
			"Access-Control-Allow-Methods": "GET,POST,PATCH,DELETE,OPTIONS",
			"Access-Control-Allow-Headers": "Content-Type",
		},
		Body: body,
	}, nil
}

func awsString(s string) *string {
	return &s
}

func main() {
	lambda.Start(handler)
}
