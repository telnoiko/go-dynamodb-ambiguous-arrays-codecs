package unmarshaling

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func TryParseSliceField(av types.AttributeValue) ([]string, error) {
	switch avTyped := av.(type) {
	case *types.AttributeValueMemberSS:
		return avTyped.Value, nil
	case *types.AttributeValueMemberS:
		return []string{avTyped.Value}, nil
	case *types.AttributeValueMemberL:
		// this is untyped json list - {"apple", "banana", 42}
		casted := make([]string, 0)
		for _, dynamoValue := range avTyped.Value {
			castedValue, ok := dynamoValue.(*types.AttributeValueMemberS)
			if !ok {
				continue
			}
			casted = append(casted, castedValue.Value)
		}

		return casted, nil
	default:
		return nil, fmt.Errorf("unsopported type of unmarshal value %v, type %T", av, av)
	}
}
