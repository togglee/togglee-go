BINARY=togglee

build: test
	@go build ./...

test-ci:
	@go generate ./...
	@go test -v ./... -cover -race -coverprofile=coverage.txt -covermode=atomic

test:
	@go generate ./...
	@go test -v -short ./...

lint:
	./bin/golangci-lint run \
        --exclude="cyclomatic complexity" \
        --exclude-use-default=false \
        --enable=revive \
        --enable=gocyclo \
        --enable=goconst \
        --enable=unconvert \
        ./...

prepare:
	@echo "Installing golangci-lint"
	# @go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s latest
	@echo "Installing gomock"
	@go get -u github.com/golang/mock/gomock
	@echo "Installing mockgen"
	@go install github.com/golang/mock/mockgen
	@echo "Installing go.mod"
	@go mod download
	@echo "Add GOPATH to PATH"
	@export PATH=$PATH:$(go env GOPATH)/bin

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: lint lint-prepare clean build unittest