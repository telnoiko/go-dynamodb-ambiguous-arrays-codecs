package repository

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	typez "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"go-dynamodb-ambiguous-arrays-codecs/app/types"
	"go-dynamodb-ambiguous-arrays-codecs/app/types/agnostic_array"
	"go-dynamodb-ambiguous-arrays-codecs/app/types/agnostic_type"
)

type DynamoRepository struct {
	dynamodbClient *dynamodb.Client
}

const UserChoicesTableName = "UserChoices"

func NewDynamoRepository(dynamoHost string) *DynamoRepository {
	dynamodbClient := createLocalClient(dynamoHost)
	return &DynamoRepository{dynamodbClient}
}

func createLocalClient(url string) *dynamodb.Client {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion("us-west-2"),
		config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(
			func(service, region string, options ...any) (aws.Endpoint, error) {
				return aws.Endpoint{URL: url}, nil
			})),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: "dummy", SecretAccessKey: "dummy", SessionToken: "dummy",
				Source: "Hard-coded credentials; values are irrelevant for local DynamoDB",
			},
		}),
	)
	if err != nil {
		panic(err)
	}

	return dynamodb.NewFromConfig(cfg)
}

func (d *DynamoRepository) SaveUserChoiceAbstract(userChoiceAbstract *types.UserChoiceRequest) (id string, err error) {
	table := UserChoicesTableName
	userChoiceAbstract.ID = uuid.New().String()

	item, err := attributevalue.MarshalMap(userChoiceAbstract)
	if err != nil {
		return "", err
	}

	request := dynamodb.PutItemInput{
		TableName: &table,
		Item:      item,
	}

	_, err = d.dynamodbClient.PutItem(context.Background(), &request)
	return userChoiceAbstract.ID, err
}

func (d *DynamoRepository) GetUserChoiceAbstract(id uuid.UUID) (*types.UserChoiceRequest, error) {
	request := mapToDynamoRequest(id)

	item, err := d.dynamodbClient.GetItem(context.Background(), &request)
	if err != nil {
		return nil, err
	}

	var userChoiceAbstract types.UserChoiceRequest
	err = attributevalue.UnmarshalMap(item.Item, &userChoiceAbstract)

	return &userChoiceAbstract, err
}

func (d *DynamoRepository) GetUserChoiceAgnosticArray(id uuid.UUID) (*agnostic_array.UserChoiceAgnosticArray, error) {
	request := mapToDynamoRequest(id)

	item, err := d.dynamodbClient.GetItem(context.Background(), &request)
	if err != nil {
		return nil, err
	}

	var userChoiceCustom agnostic_array.UserChoiceAgnosticArray
	err = attributevalue.UnmarshalMap(item.Item, &userChoiceCustom)

	return &userChoiceCustom, err
}

func (d *DynamoRepository) GetUserChoiceAgnosticType(id uuid.UUID) (*agnostic_type.UserChoiceAgnosticType, error) {
	request := mapToDynamoRequest(id)

	item, err := d.dynamodbClient.GetItem(context.Background(), &request)
	if err != nil {
		return nil, err
	}

	var userChoiceCustom agnostic_type.UserChoiceAgnosticType
	err = attributevalue.UnmarshalMap(item.Item, &userChoiceCustom)

	return &userChoiceCustom, err
}

func mapToDynamoRequest(id uuid.UUID) dynamodb.GetItemInput {
	table := UserChoicesTableName
	request := dynamodb.GetItemInput{
		TableName: &table,
		Key: map[string]typez.AttributeValue{
			"id": &typez.AttributeValueMemberS{
				Value: id.String(),
			},
		},
	}
	return request
}
