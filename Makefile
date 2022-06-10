os?=linux
port?=8686


run: 
	swag init -g http/router.go
	go run main.go bootstrap.go

build:export GOOS=$(os)
build:export GOARCH=amd64
build:
	@echo "building binary for $(GOOS)..."
	swag init -g http/router.go
	go build -o ./super_admin 
	@echo "done!"
	