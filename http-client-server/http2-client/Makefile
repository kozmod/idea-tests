build_time :=$(shell date +%Y-%m-%dT%H:%M:%S%z)

##CLiernt
client_bin :=client/http2client
client_mod_path :="./http2-client"
client_src_path :="./cmd/main/main.go"
define client_ldflags
-X 'github.com/kozmod/idea-tests/http-client-server/http2-client/version.Version=$(client_version_tag)' \
-X 'github.com/kozmod/idea-tests/http-client-server/http2-client/version.Time=$(build_time)'
endef

build.client:
	${GO_ENV} go build -o ${BIN_OUT} -ldflags="${client_ldflags}" ${client_src_path}