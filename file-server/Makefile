MAIN_PATH	:=main.go

VERSION		:=v0.0.7
REPO 		:=kozmo
TAG			:=${REPO}/file-server:${VERSION}

PORT		:=9002
FILE_STORAGE:=./files
ROOT_CERT_PATH	:=./certs/sec.syn.crt
ROOT_KEY_PATH	:=./certs/sec.syn.key
CERT_PATH		:=./certs/ott-ingress.crt
KEY_PATH		:=./certs/ott-ingress.key

run:
	GO111MODULE="on" \
	FILE_STORAGE=${FILE_STORAGE} \
	PORT=${PORT} \
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
	--build-arg app_port=${PORT} \
	--build-arg app_file_storege=${FILE_STORAGE} \
	--build-arg cert_path=${CERT_PATH} \
	--build-arg key_path=${KEY_PATH} \
	.