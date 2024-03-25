This is a demonstration purpose app to show how to use inbuilt un/marshalling
functionality of DynamoDB in AWS SDK for Golang.

## Pre-requisite

- Docker
- Golang 1.21 or higher

## How to run

- Clone the repo
- Use following [Makefile](Makefile) commands to run the app
    - `make up` to run the app and DynamoDB in docker compose locally
    - `make up-dynamo` to run only DynamoDB in docker compose locally
    - `make down` to stop the running app and/or DynamoDB

## How to debug

- 