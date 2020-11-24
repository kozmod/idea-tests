build.linux:
	GO111MODULE="on" CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/main .

build.current:
	GO111MODULE="on" CGO_ENABLED=0 go build -a -installsuffix cgo -o ./bin/main .

fmt:
	go fmt ./...

clear:
	@(rm -rf bin)