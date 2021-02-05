# Iot-Server

Using golang http package build a sample server connect with MongoDB that can create and search account. 

## Environment

Need to have Golang installed on your system, and either have MongoDB installed locally

- Go modules

```sh
go mode download
```

- MongoDB

Database name: iot

Collection name: account

## Usage

```
go run main.go
```

Entering `localhost:12345/` in your browser

## Reference

- [Enable CORS on a Go Web Server](https://flaviocopes.com/golang-enable-cors/)
