version :=v1.0.6
build_tag :=kozmod/sleep-client:${version}
tag_to_replace=registry.sigma.sbrf.ru/dev/ci01994970/ci02001129_synapse_dev/${build_tag}

docker:
	docker image build -t ${build_tag} . && make docker.rmi.builder

docker.rmi.builder:
	docker image prune --filter label=stage=builder <<< y