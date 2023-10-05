NAME=gathering_app

.PHONY: test build build-win run run-win swag

build:
	@go mod tidy
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o $(NAME) ./cmd/main.go

build-win:
	@go mod tidy
	@go build -o $(NAME) ./cmd/main.go

run: build
	@./$(NAME)

run-win: build-win
	@./$(NAME)

swag:
	swag fmt -d ./ --exclude ./internal/adapter/docs
	swag init -g ./cmd/main.go -o ./internal/adapter/docs

test:
	@go clean -testcache && go test -v -short -failfast -cover ./...
