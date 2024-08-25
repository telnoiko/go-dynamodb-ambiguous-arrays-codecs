package types

// UserDataRequest is a dto that accepts any data type for the favorite_food field
type UserDataRequest struct {
	ID           string `json:"id" dynamodbav:"id"`
	FavoriteFood any    `json:"favorite_food,omitempty" dynamodbav:"favorite_food"`
}
