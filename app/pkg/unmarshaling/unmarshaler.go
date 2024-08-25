package unmarshaling

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

func TryParseSliceField(av types.AttributeValue) ([]string, error) {
	switch av.(type) {
	case *types.AttributeValueMemberSS:
		value, _ := av.(*types.AttributeValueMemberL)

		strValues, err := readStringSlice(value)
		if err != nil {
			return nil, err
		}

		return strValues, nil
	case *types.AttributeValueMemberS:
		value, _ := av.(*types.AttributeValueMemberS)

		return []string{value.Value}, nil
	case *types.AttributeValueMemberL:
		// this is untyped json list - {"apple", "banana", 42}
		dynamoList, _ := av.(*types.AttributeValueMemberL)
		casted := make([]string, 0)

		for _, dynamoValue := range dynamoList.Value {
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
