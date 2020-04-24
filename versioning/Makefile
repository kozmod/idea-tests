version :=v1.0.0
time :=$(shell date +%Y-%m-%dT%H:%M:%S%z)
bin_output :=vers_app
main_path :=./cmd/main.go

build:
	go build -o $(bin_output) -ldflags="-X 'github.com/kozmod/idea-tests/versioning/version.VersionTime=$(time)' -X 'github.com/kozmod/idea-tests/versioning/version.Version=$(version)'" $(main_path)

vers:
	./$(bin_output) version

rm:
	rm $(bin_output)



