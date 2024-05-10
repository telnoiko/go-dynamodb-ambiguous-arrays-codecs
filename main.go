package main

import (
	"github.com/labstack/echo/v4"
	"go-dynamodb-ambiguous-arrays-codecs/app/repository"
	"go-dynamodb-ambiguous-arrays-codecs/app/usecase"
	"os"
)

func main() {
	e := echo.New()

	// this dynamodb configuration is only for local testing
	// the usecase why to change it
	// - running inside docker-compose
	// -  debugging by running app locally against dynamodb running in docker-compose
	dynamoHost := os.Getenv("DYANMODB_HOST")
	repo := repository.NewDynamoRepository(dynamoHost)
	crud := usecase.NewDataCrud(repo)

	e.POST("/choice", crud.CreateUserChoice)
	e.GET("/choice-manual/:id", crud.GetUserChoice)
	e.GET("/choice-auto-array/:id", crud.GetUserChoiceAgnosticArray)
	e.GET("/choice-auto-type/:id", crud.GetUserChoiceAgnosticType)

	e.Logger.Fatal(e.Start(":1323"))
}
