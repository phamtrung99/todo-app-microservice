## Overview

Simple go microservice todo-app using [go-micro library](https://github.com/micro/go-micro)

## Structure

```bash
todo-app-microservice/
│── api-gateway/         # HTTP Gateway using Go-Micro
│   ├── go.mod
│   ├── main.go          # Initializes the HTTP server
│   ├── handler.go       # Handles HTTP requests
│   ├── client.go        # Connects to gRPC services
│── todo/                # gRPC services
│   ├── go.mod
│   ├── main.go
│   ├── proto/
│       ├── todo.pb.proto     # Defines gRPC messages and service
│       └── todo.pb.micro.go   # File generated from todo.proto

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





