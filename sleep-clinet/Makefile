version :=v2.0.0
build_tag :=kozmo/sleep-client:${version}
tag_to_replace=registry.sigma.sbrf.ru/dev/ci01994970/ci02001129_synapse_dev/${build_tag}

fmt:
	go fmt .

docker:
	docker image build -t ${tag_to_replace} . && make docker.rmi.builder

docker.rmi.builder:
	docker image prune --filter label=stage=builder <<< y