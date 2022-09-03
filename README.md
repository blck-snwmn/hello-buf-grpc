# hello-buf-grpc

## Use protoc w/go plugins

```bash
protoc --go_out=./gen/grpc --go_opt=paths=source_relative --go-grpc_out=./gen/grpc --go-grpc_opt=paths=source_relative hello.proto
```

## Use Buf

```bash
buf generate
```
