server_port :=:8080
server_url :=http://localhost

mod_path :="github.com/kozmod/idea-tests/http-client-server"

#SERVER
server_bin :=server/http2server
server_mod_path :="./http2-server"
server_src_path :="./cmd/main/main.go"
server_version_tag :=v1.0.5

#CLiernt
client_bin :=client/http2client
client_mod_path :="./http2-client"
client_src_path :="./cmd/main/main.go"
client_version_tag :=v1.0.0

bin_dir := bin

export SERVER_ADDR=$(server_url)$(server_port)
export REQUEST_QUANTITY=10
export REQUEST_FREQUENCY_SEC=60

rebuild.run.server:
	 @(rm -rf ${bin_dir}/${server_bin} && make build.server && make run.server)
rebuild.run.client:
	 @(rm -rf ${bin_dir}/${client_bin} && make build.client && make run.client)

build.server:
	pushd ${server_mod_path} && go build -o ../${bin_dir}/${server_bin} -ldflags="-X 'github.com/kozmod/idea-tests/http-client-server/http2-server/cmd.DefaultServerPort=${server_port}'" ${server_src_path}
build.client:
	pushd ${client_mod_path} && go build -o ../${bin_dir}/${client_bin} ${client_src_path}

build.all:
	@(make clear)
	@(make build.server)
	@(make build.client)
	@(pushd ${bin_dir} && ./${server_bin} dval)
	@(pushd ${bin_dir} && ./${client_bin} dval)

run.server:
	pushd ${bin_dir} && ./${server_bin}
run.client:
	pushd ${bin_dir} && ./${client_bin}

docker.server:
	pushd ${server_mod_path} && docker image build -t go-http2-server:${server_version_tag} .
docker.client:
	pushd ${client_mod_path} && docker image build -t  go-http2-cycle-client:${client_version_tag} .

clear:
	@(rm -rf ${bin_dir})

imports:
	goimports -w .





