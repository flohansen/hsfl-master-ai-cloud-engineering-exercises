# Install Dependencies

    sudo apt install -y protobuf-compiler
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

# Compile Protocol Buffers

    protoc --go_out=. --go-grpc_out=. protobuf/hello.proto