BINARY=feaggle

build: test
	@go build ./...

test-ci: lint
	@go generate ./...
	@go test -v ./... -cover -race -coverprofile=coverage.txt -covermode=atomic

test: lint
	@go generate ./...
	@go test -v -short ./...

lint:
    ./bin/golangci-lint run \
        --exclude="cyclomatic complexity" \
        --exclude-use-default=false \
        --enable=golint \
        --enable=gocyclo \
        --enable=goconst \
        --enable=unconvert \
        ./...

prepare:
	@echo "Installing golangci-lint"
	# @go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s v1.13.2
	@echo "Installing gomock"
	@go get -u github.com/golang/mock/gomock
	@echo "Installing mockgen"
	@go install github.com/golang/mock/mockgen

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: lint lint-prepare clean build unittest