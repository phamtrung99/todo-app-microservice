ARG GOLANG_TAG=1.24.1-alpine
FROM golang:${GOLANG_TAG}
WORKDIR /app

COPY . .

RUN GOOS=linux GOARCH=arm64 go build -o api-gateway .