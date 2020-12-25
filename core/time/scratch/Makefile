version :=v0.0.0
repo := kozmo/scratch-time
build_tag :=${repo}:${version}
container_name :=scratch-time-test

docker:
	docker image build -t ${build_tag} . && make docker.rmi.builder

docker.rmi.builder:
	docker image prune --filter label=stage=builder <<< y

docker.run:
	docker run -it --name ${container_name} ${build_tag}

docker.rm:
	docker rm ${container_name} && docker images | grep ${repo} | awk '{ print $$3 }' | xargs docker rmi

download:
	go mod download

get.all:
	go get -u all

fmt:
	go fmt . && go mod tidy

