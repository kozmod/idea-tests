MAIN_PATH	:=main.go

VERSION		:=v0.0.7
REPO 		:=kozmo
TAG			:=${REPO}/file-server:${VERSION}

HTTP_PORT		:=9001
HTTPS_PORT		:=9003
FILE_STORAGE	:=./files
CERT_PATH		:=./certs/ott-ingress.crt
KEY_PATH		:=./certs/ott-ingress.key

run:
	GO111MODULE="on" \
	FILE_STORAGE=${FILE_STORAGE} \
	HTTP_PORT=${HTTP_PORT} \
	HTTPS_PORT=${HTTPS_PORT} \
	CERT_PATH=${CERT_PATH} \
	KEY_PATH=${KEY_PATH} \
	go run ${MAIN_PATH}

fmt:
	go fmt ./...
	goimports -w .

build: fmt
	go build -o bin/main ${MAIN_PATH}

docker.local: fmt
	GO111MODULE="on" CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o ./bin/main ${MAIN_PATH}
	docker image build -t ${TAG} \
	--build-arg app_http_port=${HTTP_PORT} \
	--build-arg app_https_port=${HTTPS_PORT} \
	--build-arg app_file_storege=${FILE_STORAGE} \
	--build-arg cert_path=${CERT_PATH} \
	--build-arg key_path=${KEY_PATH} \
	.