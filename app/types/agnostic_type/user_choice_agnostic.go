package agnostic_type

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"go-dynamodb-ambiguous-arrays-codecs/app/pkg/unmarshaling"
)

// UserChoiceAgnosticType is a dto with automatic type conversion
// for the whole object using type casting
type UserChoiceAgnosticType struct {
	ID     string   `json:"id" dynamodbav:"id"`
	Choice []string `json:"choice" dynamodbav:"choice"`
}

func (a *UserChoiceAgnosticType) UnmarshalDynamoDBAttributeValue(av types.AttributeValue) error {
	m, ok := av.(*types.AttributeValueMemberM)
	if !ok {
		return fmt.Errorf("failed to unmarshal UserChoiceAgnosticType from value {%v}", av)
	}

	var userChoice UserChoiceAgnosticType
	for key, value := range m.Value {
		switch key {
		case "id":
			parsed, ok := value.(*types.AttributeValueMemberS)
			if !ok {
				return fmt.Errorf("failed to unmarshal id from value {%v}, type %T", av, av)
			}

			userChoice.ID = parsed.Value
		case "choice":
			parsed, err := unmarshaling.TryParseSliceField(value)
			if err != nil {
				return err
			}

			userChoice.Choice = parsed
		}
	}

	*a = userChoice

	return nil
}
