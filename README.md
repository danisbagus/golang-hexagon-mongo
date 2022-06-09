# golang-hexagon-mongo
simple app use hexagon architecture and mongo db

## Requirements

- [Golang](https://golang.org/) as main programming language.
- [Go Module](https://go.dev/blog/using-go-modules) for package management.
- [Docker-compose](https://docs.docker.com/compose/) for running MongoDB.

## Setup

Create MongoDB container

```bash
./startdb.sh
```

## Run the service

```

Install packages

```bash
go get ./...
```

Run app

```bash
go run app/api/main.go
```
