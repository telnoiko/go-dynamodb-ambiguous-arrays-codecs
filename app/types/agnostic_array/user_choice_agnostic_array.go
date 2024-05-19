package agnostic_array

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"go-dynamodb-ambiguous-arrays-codecs/app/pkg/unmarshaling"
)

// UserChoiceAgnosticArray is a dto with automatic type conversion
// for custom string slice field
type UserChoiceAgnosticArray struct {
	ID     string        `json:"id" dynamodbav:"id"`
	Choice AgnosticArray `json:"choice" dynamodbav:"choice"`
}

type AgnosticArray []string

func (a AgnosticArray) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return &types.AttributeValueMemberSS{Value: a}, nil
}

func (a *AgnosticArray) UnmarshalDynamoDBAttributeValue(av types.AttributeValue) error {
	arr, err := unmarshaling.TryParseSliceField(av)
	if err != nil {
		return err
	}

	*a = arr

	return nil
}
