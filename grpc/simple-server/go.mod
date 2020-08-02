module github.com/kozmod/idea-tests/grpc/simple-server

go 1.13

require (
	github.com/google/go-cmp v0.5.0 // indirect
	github.com/kozmod/idea-tests/grpc/proto v0.0.0
	github.com/spf13/cobra v1.0.0
	golang.org/x/exp v0.0.0-20190121172915-509febef88a4
	google.golang.org/grpc v1.30.0
	google.golang.org/protobuf v1.25.0 // indirect
)

replace github.com/kozmod/idea-tests/grpc/proto => ../proto
