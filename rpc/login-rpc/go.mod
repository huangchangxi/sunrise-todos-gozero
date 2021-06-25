module login-rpc

go 1.15

require (
	github.com/golang/protobuf v1.5.2
	github.com/tal-tech/go-zero v1.1.7
	google.golang.org/grpc v1.36.0
	google.golang.org/grpc/examples v0.0.0-20210623211556-d9eb12feed7a // indirect
)

replace google.golang.org/grpc v1.36.0 => google.golang.org/grpc v1.29.1
