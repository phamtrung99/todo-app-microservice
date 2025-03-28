## Overview

Simple go microservice todo-app using [go-micro library](https://github.com/micro/go-micro)

## Structure

```bash
toto-app-microservice/
│── api-gateway/         # HTTP Gateway sử dụng Go-Micro
│   ├── go.mod
│   ├── main.go          # Khởi tạo HTTP server
│   ├── handler.go       # Xử lý HTTP request
│   ├── client.go        # Kết nối đến gRPC services
│── todo/                # gRPC Service 1
│   ├── go.mod
│   ├── main.go
│   ├── proto/
│── task/                # gRPC Service 2
│   ├── go.mod
│   ├── main.go
│   ├── proto/│       ├── task.proto   # Định nghĩa gRPC message và service
│       └── task.pb.go   # File được generate từ task.proto
│── proto/               # Chứa các file .proto
│   ├── todo.proto       # Định nghĩa gRPC message và service
│   └── todo.pb.go       # File được generate từ todo.proto


```


## Start
### I. Run manually: 
1. Gen protoc

``` bash
protoc --proto_path=./proto --micro_out=./proto --go_out=./proto todo.proto
```

Run RGC server
``` bash
go run todo/main.go
```

2. Run HTTP server

``` bash
go run api-gateway/main.go
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





