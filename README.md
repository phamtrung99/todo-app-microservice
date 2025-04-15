## Overview

Simple go microservice todo-app using [go-micro library](https://github.com/micro/go-micro)


## Start
### I. Local run:
**Prequisite:**  

[protoc-gen-micro](https://github.com/micro/go-micro/tree/master/cmd/protoc-gen-micro#install)  
[protoc](https://protobuf.dev/installation/)


**Gen proto file**  
File .proto was be clone from todo service to api-gateway or any client, after that generate from .proto interface to client code (JS, Go, etc...)
```
protoc --proto_path=./proto --micro_out=./proto --go_out=./proto todo.proto
```


**Start**
```
docker compose up --build
```

**Test by curl**

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

### Reference
[Gen migrate file](./todo/database/Readme.md)


