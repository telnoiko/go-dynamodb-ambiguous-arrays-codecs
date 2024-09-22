package manual

import "fmt"

type UserDataTarget struct {
	FavoriteFood []string `json:"favorite_food"`
}

// ConvertToArray is a manual conversion function to map 2 possible
// types of fields in DynamoDB - string and array of strings -
// into a slice of strings.
func ConvertToArray(field any) ([]string, error) {
	if field == nil {
		return nil, nil
	}

	switch field.(type) {
	case []string:
		return field.([]string), nil
	case string:
		return []string{field.(string)}, nil
	case []any:
		values := field.([]any)
		casted := make([]string, 0)
		for _, value := range values {
			castedValue, ok := value.(string)
			if !ok {
				continue
			}
			casted = append(casted, castedValue)
		}

		return casted, nil
	default:
		return nil, fmt.Errorf("unsupported type '%T' for the field '%v'", field, field)
	}
}
