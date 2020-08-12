# define standard colors
BLACK        := $(shell tput -Txterm setaf 0)
RED          := $(shell tput -Txterm setaf 1)
GREEN        := $(shell tput -Txterm setaf 2)
YELLOW       := $(shell tput -Txterm setaf 3)
LIGHTPURPLE  := $(shell tput -Txterm setaf 4)
PURPLE       := $(shell tput -Txterm setaf 5)
BLUE         := $(shell tput -Txterm setaf 6)
WHITE        := $(shell tput -Txterm setaf 7)

# reset colours after message
RESET := $(shell tput -Txterm sgr0)

#param
src_path :=./
version_tag :=v1.0.5
project_name :=k8s-test-metrics-reader

docker:
	@echo "\n> 🚀 ${GREEN}  Building ${project_name} ${RESET}> \n"
	@(pushd ${src_path} \
	&& docker image build -t ${project_name}:${version_tag} . \
	&& popd)
	@(make docker.rmi.builder)

docker.rmi.builder:
	docker image prune --filter label=stage=builder <<< y

minikube.cache:
	minikube cache add ${project_name}:${version_tag}

metrics-server:
	kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/download/v0.3.7/components.yaml