package types

// UserChoiceRequest is a dto that accepts any data type for the choice field
type UserChoiceRequest struct {
	ID     string `json:"id" dynamodbav:"id"`
	Choice any    `json:"choice,omitempty" dynamodbav:"choice"`
}

// UserChoiceResponse is a dto for response body with manual type conversion
type UserChoiceResponse struct {
	Choice []string `json:"choice"`
}
