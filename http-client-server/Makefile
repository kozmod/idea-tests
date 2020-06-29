build_time :=$(shell date +%Y-%m-%dT%H:%M:%S%z)

#SERVER
server_port :=:8080
server_url :=http://localhost
server_bin :=server/http2server
server_mod_path :="./http2-server"
server_src_path :="./cmd/main/main.go"
server_config_dir :=etc/config/
server_version_tag :=v1.0.9
define server_ldflags
-X 'github.com/kozmod/idea-tests/http-client-server/http2-server/cmd.DefaultServerPort=${server_port}' \
-X 'github.com/kozmod/idea-tests/http-client-server/http2-server/version.Version=$(server_version_tag)' \
-X 'github.com/kozmod/idea-tests/http-client-server/http2-server/version.Time=$(build_time)'
endef

#CLiernt
client_bin :=client/http2client
client_mod_path :="./http2-client"
client_src_path :="./cmd/main/main.go"
client_version_tag :=v1.0.1
define client_ldflags
-X 'github.com/kozmod/idea-tests/http-client-server/http2-client/version.Version=$(client_version_tag)' \
-X 'github.com/kozmod/idea-tests/http-client-server/http2-client/version.Time=$(build_time)'
endef

bin_dir := bin

#env to run client
export SERVER_ADDR=$(server_url)$(server_port)
export REQUEST_QUANTITY=10
export REQUEST_FREQUENCY_SEC=66
export POST_WITH_PAYLOAD_URL="/tp"

rebuild.run.server:
	 @(rm -rf ${bin_dir}/${server_bin} && make build.server && make run.server)
rebuild.run.client:
	 @(rm -rf ${bin_dir}/${client_bin} && make build.client && make run.client)

build.server:
	pushd ${server_mod_path} \
	&& go build -o ../${bin_dir}/${server_bin} -ldflags="${server_ldflags}" ${server_src_path} \
	&& mkdir -p ../${bin_dir}/${server_config_dir} \
	&& cp -a ./${server_config_dir}/config.yml ../${bin_dir}/${server_config_dir}
build.client:
	pushd ${client_mod_path} \
	&& go build -o ../${bin_dir}/${client_bin} -ldflags="${client_ldflags}" ${client_src_path}

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
	@(pushd ${server_mod_path} \
	&& docker image build -t go-http2-server:${server_version_tag} . \
	&& popd \
	&& make docker.rmi.builder)
docker.client:
	@(pushd ${client_mod_path} \
	&& docker image build -t go-http2-cycle-client:${client_version_tag} . \
	&& popd \
	&& make docker.rmi.builder)

docker.rmi.builder:
	docker image prune --filter label=stage=builder <<< y

clear:
	@(rm -rf ${bin_dir})

imports:
	goimports -w .





