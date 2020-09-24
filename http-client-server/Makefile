include http2-server/Makefile
include http2-client/Makefile

build_time :=$(shell date +%Y-%m-%dT%H:%M:%S%z)
bin_dir := bin

##SERVER
server_version_tag :=v1.0.15
server_bin :=server/http2server
server_mod_path :="./http2-server"
server_config_dir :=etc/config/

##CLiernt
client_version_tag :=v1.0.7
client_bin :=client/http2client
client_mod_path :="./http2-client"
server_url :=http://localhost

##env to run client
export SERVER_ADDR=$(server_url):$(server_port)
export REQUEST_QUANTITY=10
export REQUEST_FREQUENCY_SEC=66
export POST_WITH_PAYLOAD_URL=/ping

#go env
GO_ENV := GO111MODULE="on"

build.server.local:
	pushd ${server_mod_path} \
	&& make build.server BIN_OUT=../${bin_dir}/${server_bin}\
	&& mkdir -p ../${bin_dir}/${server_config_dir} \
	&& cp -a ./${server_config_dir}/config.yml ../${bin_dir}/${server_config_dir}
build.client.local:
	pushd ${client_mod_path} \
	&& make build.client BIN_OUT=../${bin_dir}/${client_bin}

run.server.local:
	pushd ${bin_dir} && ./${server_bin}
run.client.local:
	pushd ${bin_dir} && ./${client_bin}

check.server.local:
	pushd ${bin_dir} && ./${server_bin} dval && ./${server_bin} version
check.client.local:
	pushd ${bin_dir} && ./${client_bin} dval && ./${client_bin} version

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