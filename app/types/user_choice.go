package types

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// UserChoiceRequest is a dto that accepts any data type for the choice field
type UserChoiceRequest struct {
	ID     string `json:"id" dynamodbav:"id"`
	Choice any    `json:"choice,omitempty" dynamodbav:"choice"`
}

// UserChoiceResponse is a dto for response body with manual type conversion
type UserChoiceResponse struct {
	Choice []string `json:"choice"`
}

// UserChoiceCustomResponse is a dto for response body with automatic type conversion
type UserChoiceCustomResponse struct {
	ID     string        `json:"id" dynamodbav:"id"`
	Choice AgnosticArray `json:"choice" dynamodbav:"choice"`
}

type AgnosticArray []string

func (a AgnosticArray) MarshalDynamoDBAttributeValue() (types.AttributeValue, error) {
	return &types.AttributeValueMemberSS{Value: a}, nil
}

func (a *AgnosticArray) UnmarshalDynamoDBAttributeValue(av types.AttributeValue) error {
	arr, ok := av.(*types.AttributeValueMemberL)
	if !ok {
		value, ok := av.(*types.AttributeValueMemberS)
		if !ok {
			return fmt.Errorf("failed to unmarshal array from value {%v}", av)
		}
		*a = AgnosticArray{value.Value}

		return nil
	}

	var strValues []string
	for _, value := range arr.Value {
		strValue, ok := value.(*types.AttributeValueMemberS)
		if !ok {
			return fmt.Errorf("cannot parse '%v' into string", value)
		}
		strValues = append(strValues, strValue.Value)
	}

	*a = strValues

	return nil
}
