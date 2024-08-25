package main

import (
	"github.com/labstack/echo/v4"
	"go-dynamodb-ambiguous-arrays-codecs/app/repository"
	"go-dynamodb-ambiguous-arrays-codecs/app/usecase"
	"os"
)

func main() {
	e := echo.New()

	// this dynamodb configuration is only for testing local use cases
	dynamoHost := os.Getenv("DYANMODB_HOST")
	repo := repository.NewDynamoRepository(dynamoHost)
	ucase := usecase.New(repo)

	e.POST("/user-data", ucase.CreateUserData)
	e.GET("/user-data-manual/:id", ucase.GetUserData)
	e.GET("/user-data-auto-array/:id", ucase.GetUserDataAgnosticArray)
	e.GET("/user-data-auto-type/:id", ucase.GetUserDataAgnosticType)
	e.GET("/user-data-auto-reflection/:id", ucase.GetUserDataAgnosticTypeReflection)

	e.Logger.Fatal(e.Start(":1323"))
}
