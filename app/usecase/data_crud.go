package usecase

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"go-dynamodb-ambiguous-arrays-codecs/app/types"
	"go-dynamodb-ambiguous-arrays-codecs/app/types/dynamo_agnostic_slice"
	"go-dynamodb-ambiguous-arrays-codecs/app/types/dynamo_agnostic_type"
	"go-dynamodb-ambiguous-arrays-codecs/app/types/dynamo_reflection"
	"go-dynamodb-ambiguous-arrays-codecs/app/types/manual"
	"net/http"
)

type Usecase struct {
	repo userDataRepository
}

type userDataRepository interface {
	SaveUserDataAbstract(UserDataAbstract *types.UserDataRequest) (id string, err error)
	GetUserDataAbstract(id uuid.UUID) (*types.UserDataRequest, error)
	GetUserDataAgnosticArray(id uuid.UUID) (*dynamo_agnostic_slice.UserDataAgnosticSlice, error)
	GetUserDataAgnosticType(id uuid.UUID) (*dynamo_agnostic_type.UserDataAgnosticType, error)
	GetUserDataAgnosticTypeReflection(id uuid.UUID) (*dynamo_reflection.UserDataAgnosticTypeReflection, error)
}

func New(repo userDataRepository) *Usecase {
	return &Usecase{repo: repo}
}

// CreateUserData accepts any type of 'favorite_food' field  and saves it to the database
func (d *Usecase) CreateUserData(ctx echo.Context) error {
	ctx.Logger().Info("CreateUserData")
	var request types.UserDataRequest
	err := ctx.Bind(&request)
	if err != nil {
		return ctx.String(http.StatusBadRequest, fmt.Sprintf("couldn't parse body: %v", err))
	}

	id, err := d.repo.SaveUserDataAbstract(&request)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, fmt.Sprintf("couldn't save item: %v", err))
	}

	return ctx.JSON(http.StatusOK, types.UserDataRequest{ID: id})
}

// GetUserData parses the 'favorite_food' field with manual parser
func (d *Usecase) GetUserData(ctx echo.Context) error {
	ctx.Logger().Info("GetUserData")
	id := ctx.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		return ctx.String(http.StatusBadRequest, fmt.Sprintf("couldn't parse id: %v", err))
	}

	rawUserData, err := d.repo.GetUserDataAbstract(uid)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, fmt.Sprintf("couldn't get item: %v", err))
	}

	// Convert manually
	userData, err := parseAbstractRequest(rawUserData)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, fmt.Sprintf("couldn't parse data: %v", err))
	}

	return ctx.JSON(http.StatusOK, userData)
}

func parseAbstractRequest(request *types.UserDataRequest) (manual.UserDataTarget, error) {
	target := manual.UserDataTarget{}
	parsed, err := manual.ConvertToArray(request.FavoriteFood)
	if err != nil {
		return target, err
	}
	target.FavoriteFood = parsed
	return target, nil
}

// GetUserDataAgnosticArray parses the 'favorite_food' field with 'AgnosticSlice' type
// that implements inbuilt dynamodb 'Unmarshaler' interface
func (d *Usecase) GetUserDataAgnosticArray(ctx echo.Context) error {
	ctx.Logger().Info("GetUserDataAgnosticArray")
	id := ctx.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		return ctx.String(http.StatusBadRequest, fmt.Sprintf("couldn't parse id: %v", err))
	}

	UserData, err := d.repo.GetUserDataAgnosticArray(uid)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, fmt.Sprintf("couldn't get item: %v", err))
	}

	return ctx.JSON(http.StatusOK, UserData)
}

// GetUserDataAgnosticType parses the 'favorite_food' field inside a structure 'UserDataAgnosticType'
// that implements inbuilt dynamodb 'Unmarshaler' interface
func (d *Usecase) GetUserDataAgnosticType(ctx echo.Context) error {
	ctx.Logger().Info("GetUserDataAgnosticType")
	id := ctx.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		return ctx.String(http.StatusBadRequest, fmt.Sprintf("couldn't parse id: %v", err))
	}

	UserData, err := d.repo.GetUserDataAgnosticType(uid)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, fmt.Sprintf("couldn't get item: %v", err))
	}

	return ctx.JSON(http.StatusOK, UserData)
}

// GetUserDataAgnosticTypeReflection parses the 'favorite_food' field inside a structure 'UserDataAgnosticTypeReflection'
// that implements inbuilt dynamodb 'Unmarshaler' interface using dynamo_reflection
func (d *Usecase) GetUserDataAgnosticTypeReflection(ctx echo.Context) error {
	ctx.Logger().Info("GetUserDataAgnosticType")
	id := ctx.Param("id")
	uid, err := uuid.Parse(id)
	if err != nil {
		return ctx.String(http.StatusBadRequest, fmt.Sprintf("couldn't parse id: %v", err))
	}

	UserData, err := d.repo.GetUserDataAgnosticTypeReflection(uid)
	if err != nil {
		return ctx.String(http.StatusInternalServerError, fmt.Sprintf("couldn't get item: %v", err))
	}

	return ctx.JSON(http.StatusOK, UserData)
}
