build_time :=$(shell date +%Y-%m-%dT%H:%M:%S%z)

##SERVER
server_port :=8080
server_url :=http://localhost
server_src_path :="./cmd/main/main.go"
define server_ldflags
-X 'github.com/kozmod/idea-tests/http-client-server/http2-server/cmd.DefaultServerPort=${server_port}' \
-X 'github.com/kozmod/idea-tests/http-client-server/http2-server/version.Version=$(server_version_tag)' \
-X 'github.com/kozmod/idea-tests/http-client-server/http2-server/version.Time=$(build_time)'
endef

build.server:
	${GO_ENV} go build -o ${BIN_OUT} -ldflags="${server_ldflags}" ${server_src_path}
