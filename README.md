# gRPC-quickStart
Through this project you will quickly learn how to use GRPC and understand its core principles


## 安装环境

- 安装 `protoc`:

```bash
brew install protobuf
```

- 安装 `protoc-gen-go` 和 `protoc-gen-go-grpc`

```bash
go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

- 安装 `protoc-gen-grpc-gateway` 和 `protoc-gen-openapiv2`

```bash
go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway
go get github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2
