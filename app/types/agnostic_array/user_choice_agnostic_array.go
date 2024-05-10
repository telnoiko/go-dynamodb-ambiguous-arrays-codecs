package agnostic_array

import (
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// UserChoiceAgnosticArray is a dto for response body with automatic type conversion
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
	arr, err := TryParseSliceField(av)
	if err != nil {
		return err
	}

	*a = arr

	return nil
}

func TryParseSliceField(av types.AttributeValue) ([]string, error) {
	switch av.(type) {
	case *types.AttributeValueMemberL:
		value, _ := av.(*types.AttributeValueMemberL)

		strValues, err := readStringSlice(value)
		if err != nil {
			return nil, err
		}

		return strValues, nil
	case *types.AttributeValueMemberS:
		value, _ := av.(*types.AttributeValueMemberS)

		return []string{value.Value}, nil
	default:
		return nil, fmt.Errorf("unsopported type of unmarshal value %v, type %T", av, av)
	}
}

func readStringSlice(av *types.AttributeValueMemberL) ([]string, error) {
	var strValues []string
	for _, value := range av.Value {
		strValue, ok := value.(*types.AttributeValueMemberS)
		if !ok {
			return nil, fmt.Errorf("cannot parse '%v' into string", value)
		}
		strValues = append(strValues, strValue.Value)
	}
	return strValues, nil
}
