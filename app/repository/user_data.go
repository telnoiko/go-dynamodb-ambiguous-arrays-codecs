package repository

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	typez "github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/google/uuid"
	"go-dynamodb-ambiguous-arrays-codecs/app/types"
	"go-dynamodb-ambiguous-arrays-codecs/app/types/dynamo_agnostic_array"
	"go-dynamodb-ambiguous-arrays-codecs/app/types/dynamo_agnostic_type"
	"go-dynamodb-ambiguous-arrays-codecs/app/types/dynamo_reflection"
)

func (d *DynamoRepository) SaveUserDataAbstract(UserDataAbstract *types.UserDataRequest) (id string, err error) {
	UserDataAbstract.ID = uuid.New().String()

	item, err := attributevalue.MarshalMap(UserDataAbstract)
	if err != nil {
		return "", err
	}

	table := UserDataTableName
	request := dynamodb.PutItemInput{
		TableName: &table,
		Item:      item,
	}

	_, err = d.dynamodbClient.PutItem(context.Background(), &request)
	return UserDataAbstract.ID, err
}

func (d *DynamoRepository) GetUserDataAbstract(id uuid.UUID) (*types.UserDataRequest, error) {
	request := mapToDynamoRequest(id)

	item, err := d.dynamodbClient.GetItem(context.Background(), &request)
	if err != nil {
		return nil, err
	}

	var UserDataAbstract types.UserDataRequest
	err = attributevalue.UnmarshalMap(item.Item, &UserDataAbstract)

	return &UserDataAbstract, err
}

func (d *DynamoRepository) GetUserDataAgnosticArray(id uuid.UUID) (*dynamo_agnostic_array.UserDataAgnosticArray, error) {
	request := mapToDynamoRequest(id)

	item, err := d.dynamodbClient.GetItem(context.Background(), &request)
	if err != nil {
		return nil, err
	}

	var UserDataCustom dynamo_agnostic_array.UserDataAgnosticArray
	err = attributevalue.UnmarshalMap(item.Item, &UserDataCustom)

	return &UserDataCustom, err
}

func (d *DynamoRepository) GetUserDataAgnosticType(id uuid.UUID) (*dynamo_agnostic_type.UserDataAgnosticType, error) {
	request := mapToDynamoRequest(id)

	item, err := d.dynamodbClient.GetItem(context.Background(), &request)
	if err != nil {
		return nil, err
	}

	var UserDataCustom dynamo_agnostic_type.UserDataAgnosticType
	err = attributevalue.UnmarshalMap(item.Item, &UserDataCustom)

	return &UserDataCustom, err
}

func (d *DynamoRepository) GetUserDataAgnosticTypeReflection(id uuid.UUID) (*dynamo_reflection.UserDataAgnosticTypeReflection, error) {
	request := mapToDynamoRequest(id)

	item, err := d.dynamodbClient.GetItem(context.Background(), &request)
	if err != nil {
		return nil, err
	}

	var UserDataCustom dynamo_reflection.UserDataAgnosticTypeReflection
	err = attributevalue.UnmarshalMap(item.Item, &UserDataCustom)

	return &UserDataCustom, err
}

func mapToDynamoRequest(id uuid.UUID) dynamodb.GetItemInput {
	table := UserDataTableName
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
