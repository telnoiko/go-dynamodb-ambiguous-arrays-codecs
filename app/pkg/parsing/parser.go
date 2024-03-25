package parsing

import (
	"fmt"
)

// ConvertToArray is a manual conversion function to map 2 possible
// types of fields in dynamodb - string and array of strings -
// into a slice of strings.
func ConvertToArray(field any) []string {
	array, err := tryConvertIfArray(field)
	if err != nil {
		value, err := tryConvertIfString(field)
		if err != nil {
			return nil
		}
		return []string{*value}
	}

	return array
}

func tryConvertIfString(field any) (*string, error) {
	value, ok := field.(string)
	if !ok {
		return nil, fmt.Errorf("cannot parse '%v' into string", field)
	}

	return &value, nil
}

func tryConvertIfArray(field any) ([]string, error) {
	values, ok := field.([]any)
	if !ok {
		return nil, fmt.Errorf("cannot parse value '%v' into array of strings", field)
	}

	var strValues []string
	for _, value := range values {
		strValue, ok := value.(string)
		if !ok {
			return nil, fmt.Errorf("cannot parse '%v' into string", value)
		}
		strValues = append(strValues, strValue)
	}

	return strValues, nil
}
