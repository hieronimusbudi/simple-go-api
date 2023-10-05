# Gathering REST API

Simple API using (Golang, Gin, MySQL and Docker).

## Documentation

This API documentation using Swagger, please build and run app first to access Swagger URL

```
http://localhost:3000/swagger/docs/index.html#/
```

Use make command to generate swagger files, in docs folder

```
go install github.com/swaggo/swag/cmd/swag@latest //please install swag first

make swag
```

## How to run

### Using Docker Compose

Use this command to run docker compose. Please change db host in .env into `DBHOST=mysql-db:3306`

```
docker compose -f docker.compose.yml up --build
```

### Local

Use `make` command to run or build app. Please change db host in .env into `DBHOST=localhost:3306`

```
make build
make run //for linux
make run-win //for windows
```

## Testing

There are 2 testing types, unit test for mostly code and integration test for adapter repository code. Integration test using Docker to create test DB.

### Run test

```
make test
```
