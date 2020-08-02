module github.com/kozmod/idea-tests/grpc/simple-server

go 1.13

require (
	github.com/golang/protobuf v1.4.2
	github.com/google/go-cmp v0.5.0 // indirect
	github.com/kozmod/idea-tests/grpc/proto v0.0.0
	github.com/spf13/cobra v1.0.0
	google.golang.org/grpc v1.30.0
	google.golang.org/protobuf v1.23.0
)

replace github.com/kozmod/idea-tests/grpc/proto => ../proto
