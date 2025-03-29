ARG GOLANG_TAG=1.24.1-alpine
FROM golang:${GOLANG_TAG}
WORKDIR /home/api-gateway

COPY . .

CMD ["go","run","main.go"]