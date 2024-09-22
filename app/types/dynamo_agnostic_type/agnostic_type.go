package dynamo_agnostic_type

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"go-dynamodb-ambiguous-arrays-codecs/app/pkg/unmarshaling"
)

// UserDataAgnosticType is a dto with automatic type conversion
// for the whole object using type casting
type UserDataAgnosticType struct {
	ID           string   `json:"id" dynamodbav:"id"`
	FavoriteFood []string `json:"favorite_food" dynamodbav:"favorite_food"`
}

func (a *UserDataAgnosticType) UnmarshalDynamoDBAttributeValue(av types.AttributeValue) error {
	m, ok := av.(*types.AttributeValueMemberM)
	if !ok {
		return fmt.Errorf("failed to unmarshal UserDataAgnosticType from value {%v}", av)
	}

	var UserData UserDataAgnosticType
	for key, value := range m.Value {
		switch key {
		case "id":
			parsed, ok := value.(*types.AttributeValueMemberS)
			if !ok {
				return fmt.Errorf("failed to unmarshal id from value {%v}, type %T", av, av)
			}

			UserData.ID = parsed.Value
		case "favorite_food":
			parsed, err := unmarshaling.TryParseSliceField(value)
			if err != nil {
				return err
			}

			UserData.FavoriteFood = parsed
		}
	}

	*a = UserData

	return nil
}
