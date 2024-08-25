package dynamo_reflection

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"go-dynamodb-ambiguous-arrays-codecs/app/pkg/unmarshaling"
	"reflect"
)

// UserDataAgnosticTypeReflection is a dto with automatic type conversion
// for the whole object using dynamo_reflection
type UserDataAgnosticTypeReflection struct {
	ID           string   `json:"id" dynamodbav:"id"`
	FavoriteFood []string `json:"favorite_food" dynamodbav:"favorite_food"`
}

func (a *UserDataAgnosticTypeReflection) UnmarshalDynamoDBAttributeValue(av types.AttributeValue) error {
	var parsed UserDataAgnosticTypeReflection

	dynamoMap, ok := av.(*types.AttributeValueMemberM)
	if !ok {
		return fmt.Errorf("failed to unmarshal UserDataAgnosticType from value {%v}", av)
	}

	// iterate over all fields of the dto,
	// try to find matching value for it in the dynamodb object and parse it
	for _, v := range reflect.VisibleFields(reflect.TypeOf(parsed)) {
		// check if field can be parsed as dynamo type
		tag := v.Tag.Get("dynamodbav")
		if tag == "" {
			continue
		}

		// check if field is present in dynamodb object
		var dbValue, ok = dynamoMap.Value[tag]
		if !ok {
			continue
		}

		// parse the value from the db depending on the target dto type
		fields := reflect.ValueOf(&parsed).Elem()
		switch v.Type.Kind() {
		case reflect.String:
			parsed, ok := dbValue.(*types.AttributeValueMemberS)
			if !ok {
				return fmt.Errorf("failed to unmarshal id from value {%v}, type %T", av, av)
			}
			fields.FieldByName(v.Name).SetString(parsed.Value)
		case reflect.Slice:
			parsed, err := unmarshaling.TryParseSliceField(dbValue)
			if err != nil {
				return err
			}
			fields.FieldByName(v.Name).Set(reflect.ValueOf(parsed))
		default:
			return fmt.Errorf("unsupported dto type {%T}, value {%v}", v.Type.Kind(), v)
		}
	}

	*a = parsed

	return nil
}
