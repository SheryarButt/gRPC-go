# Learning gRPC using Golang

- Generate code through protoc for gRPC manually using:
  - protoc -Igreet/proto --go_out=. --go-grpc_out=. greet/proto/dummy.proto (Without go mod)
  - protoc -Igreet/proto --go_out=. --go_opt=module=github.com/SheryarButt/gRPC-go --go-grpc_out=. --go-grpc_opt=module=github.com/SheryarButt/gRPC-go greet/proto/dummy.proto (With go mod)
  - Specific to this code only.
