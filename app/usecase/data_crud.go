package usecase

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go-dynamodb-ambiguous-arrays-codecs/app/pkg/parsing"
	"go-dynamodb-ambiguous-arrays-codecs/app/repository"
	"go-dynamodb-ambiguous-arrays-codecs/app/types"
	"net/http"
)

type DataCrud struct {
	repo repository.DynamoRepository
}

func NewDataCrud(repo *repository.DynamoRepository) *DataCrud {
	return &DataCrud{repo: *repo}
}

// CreateUserChoice accepts any type of 'choice' field  and saves it to the database
func (d *DataCrud) CreateUserChoice(ctx echo.Context) error {
	ctx.Logger().Info("CreateUserChoice")
	var choice types.UserChoiceRequest
	err := ctx.Bind(&choice)
	if err != nil {
		return ctx.String(http.StatusBadRequest, fmt.Sprintf("couldn't parse body: %v", err))
	}

	id, err := d.repo.SaveUserChoiceAbstract(&choice)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, fmt.Sprintf("couldn't save item: %v", err))
	}

	return ctx.JSON(http.StatusOK, types.UserChoiceRequest{ID: id})
}

// GetUserChoice parses the 'choice' field with manual parser
func (d *DataCrud) GetUserChoice(ctx echo.Context) error {
	ctx.Logger().Info("GetUserChoice")
	id := ctx.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		return ctx.String(http.StatusBadRequest, fmt.Sprintf("couldn't parse id: %v", err))
	}

	rawChoice, err := d.repo.GetUserChoiceAbstract(uid)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, fmt.Sprintf("couldn't get item: %v", err))
	}

	// Convert manually
	choice := parseAbstractChoice(rawChoice)

	return ctx.JSON(http.StatusOK, choice)
}

func parseAbstractChoice(rawChoice *types.UserChoiceRequest) types.UserChoiceResponse {
	choice := types.UserChoiceResponse{}
	choiceField := parsing.ConvertToArray(rawChoice.Choice)
	choice.Choice = choiceField
	return choice
}

// GetUserChoiceAgnosticArray parses the 'choice' field with 'AgnosticArray' type
// that implements inbuilt dynamodb 'Unmarshaler' interface
func (d *DataCrud) GetUserChoiceAgnosticArray(ctx echo.Context) error {
	ctx.Logger().Info("GetUserChoiceAgnosticArray")
	id := ctx.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		return ctx.String(http.StatusBadRequest, fmt.Sprintf("couldn't parse id: %v", err))
	}

	userChoice, err := d.repo.GetUserChoiceAgnosticArray(uid)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, fmt.Sprintf("couldn't get item: %v", err))
	}

	return ctx.JSON(http.StatusOK, userChoice)
}

// GetUserChoiceAgnosticType parses the 'choice' field inside a structure 'UserChoiceAgnosticType'
// that implements inbuilt dynamodb 'Unmarshaler' interface
func (d *DataCrud) GetUserChoiceAgnosticType(ctx echo.Context) error {
	ctx.Logger().Info("GetUserChoiceAgnosticType")
	id := ctx.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		return ctx.String(http.StatusBadRequest, fmt.Sprintf("couldn't parse id: %v", err))
	}

	userChoice, err := d.repo.GetUserChoiceAgnosticType(uid)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, fmt.Sprintf("couldn't get item: %v", err))
	}

	return ctx.JSON(http.StatusOK, userChoice)
}

// GetUserChoiceAgnosticTypeReflection parses the 'choice' field inside a structure 'UserChoiceAgnosticTypeReflection'
// that implements inbuilt dynamodb 'Unmarshaler' interface using reflection
func (d *DataCrud) GetUserChoiceAgnosticTypeReflection(ctx echo.Context) error {
	ctx.Logger().Info("GetUserChoiceAgnosticType")
	id := ctx.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		return ctx.String(http.StatusBadRequest, fmt.Sprintf("couldn't parse id: %v", err))
	}

	userChoice, err := d.repo.GetUserChoiceAgnosticTypeReflection(uid)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, fmt.Sprintf("couldn't get item: %v", err))
	}

	return ctx.JSON(http.StatusOK, userChoice)
}
