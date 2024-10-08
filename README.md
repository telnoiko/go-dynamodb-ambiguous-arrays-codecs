This is a demonstration purpose app to show how to use inbuilt un/marshalling
functionality of DynamoDB in AWS SDK for Golang.

## Pre-requisite

- Docker
- Golang 1.22 or higher

## How to run

- Clone the repo
- Use following [Makefile](Makefile) commands to run the app
    - `make up` to run the app and DynamoDB in Docker locally, or
    - `make up-dynamo` to run DynamoDB in Docker locally
    - `make run` to run the application against running DynamoDB locally
- After you're done shutdown containers by running `make down`

## How to debug

- Run DynamoDB in container by running `make up-dynamo`
- Install dependencies by running `make deps`
- Run the app in debug mode from yor IDE of choice and passing environment
  variable to the container `DYANMODB_HOST=http://localhost:8000`
- After you're done shutdown containers by running `make down`