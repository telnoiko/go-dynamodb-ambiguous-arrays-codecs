package unmarshaling

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

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
