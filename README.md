## Overview

Simple go microservice todo-app using [go-micro library](https://github.com/micro/go-micro)

## Structure

```
└── todo-app-microservice
    └── api-gateway
        └── client
            └── client.go
        └── Dockerfile
        └── go.mod
        └── go.sum
        └── handler
            └── handler.go
        └── main.go
    └── auth
    └── todo
        └── Dockerfile
        └── go.mod
        └── go.sum
        └── main.go
        └── proto
            └── todo.pb.go
            └── todo.pb.micro.go
            └── todo.proto
    └── docker-compose.yml
    └── README.md
```


## Start
### I. Run by docker compose: 

```
docker compose up --build
```

1. Gen protoc

``` bash
protoc --proto_path=./proto --micro_out=./proto --go_out=./proto todo.proto
```

3. Test by curl

Create Todo
``` bash
curl -X POST "http://localhost:3000/todos" \
     -H "Content-Type: application/json" \
     -d '{"title": "Learn Go-Micro", "description": "Build a microservice using Go-Micro"}'
```

Get Todos
``` bash
curl -X GET "http://localhost:3000/todos"
```





