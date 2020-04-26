server_port :=":8080"
server_url :="http://localhost"
server_addr :=$(server_url)$(server_port)

mod_path :="github.com/kozmod/idea-tests/http-client-server"

server_bin_output :=http2server
server_mod_path :="./http2-server"
server_src_path :="./cmd/main/main.go"
server_version_tag :=v1.0.0

client_bin_output :=http2client
client_mod_path :="./http2-client"
client_src_path :="./cmd/main/main.go"
client_version_tag :=v1.0.0

bin_dir :="../bin"

build.server:
	pushd $(server_mod_path) && go build -o $(bin_dir)/$(server_bin_output) -ldflags="-X 'github.com/kozmod/idea-tests/http-client-server/http2-server/cmd.DefaultServerPort=$(server_port)'" $(server_src_path)
build.client:
	pushd $(client_mod_path) && go build -o $(bin_dir)/$(client_bin_output) -ldflags="-X 'github.com/kozmod/idea-tests/http-client-server/http2-client/cmd.DefaultServerAddr=$(server_addr)'" $(client_src_path)

build.all:
	@([ -d "./bin" ] && rm -r ./bin)
	make build.server && ./bin/$(server_bin_output) dval
	make build.client && ./bin/$(client_bin_output) dval

run.server:
	pushd bin && ./$(server_bin_output)
run.client:
	pushd bin && ./$(client_bin_output)

docker.server:
	pushd $(server_mod_path) && docker image build -t go-http2-server:$(server_version_tag) .
docker.client:
	pushd $(client_mod_path) && docker image build -t go-http2-client:$(client_version_tag) .

clear:
	@([ -d "./bin" ] && rm -r ./bin)





