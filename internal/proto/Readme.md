- Install Go packages `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest` and
- Install `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest` 

- Install `protoc` compiler - `brew install protobuf`
- Run `protoc --go_out=. --go-grpc_out=. url.proto` in `/internal/proto/` directory

Now you can see two newly generated files after running `protoc` command

- `url_grpc.pb.go` and `url.pb.go`
