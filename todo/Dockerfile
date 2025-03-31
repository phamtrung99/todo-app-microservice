ARG GOLANG_TAG=1.24.1-alpine
FROM golang:${GOLANG_TAG}

WORKDIR /app

ENV PATH="$PATH:$(go env GOPATH)/bin"

COPY . .

RUN set -x \
    && apk update --no-cache \
    && apk add --no-cache \
           protoc \
           protobuf-dev \
           git \
           build-base \
    && go install github.com/micro/go-micro/cmd/protoc-gen-micro \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@latest \
    && go mod tidy
    
# gen proto file 
RUN protoc --proto_path=./proto --micro_out=./proto --go_out=./proto todo.proto

# go build
RUN GOOS=linux GOARCH=arm64 go build -o todo .

