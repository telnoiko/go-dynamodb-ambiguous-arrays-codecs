package dynamo_agnostic_slice

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"go-dynamodb-ambiguous-arrays-codecs/app/pkg/unmarshaling"
)

// UserDataAgnosticSlice is a dto with automatic type conversion
// for custom string slice field
type UserDataAgnosticSlice struct {
	ID           string        `json:"id" dynamodbav:"id"`
	FavoriteFood AgnosticSlice `json:"favorite_food" dynamodbav:"favorite_food"`
}

type AgnosticSlice []string

func (a *AgnosticSlice) UnmarshalDynamoDBAttributeValue(av types.AttributeValue) error {
	arr, err := unmarshaling.TryParseSliceField(av)
	if err != nil {
		return err
	}

	*a = arr

	return nil
}
