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
│   ├── proto/

```


## Start
### I. Run manually: 
1. Gen protoc

``` bash
protoc --proto_path=./proto --micro_out=./proto --go_out=./proto todo.proto
```




### Reference
https://www.twilio.com/en-us/blog/a-practical-guide-to-creating-microservices-with-go-micro