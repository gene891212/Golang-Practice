# Iot-Server

## Environment

Need to have Golang installed on your system, and either have MongoDB installed locally

- Go modules

```sh
go get go.mongodb.org/mongo-driver/mongo
```

- MongoDB

Database name: iot

Collection name: account

## Usage

```
go run main.go
```

Entering `localhost:9090` to get all the user,
 and entering `localhost:9090/create?account=<account>&password=<password>` to create an account