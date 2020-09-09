docker:
	docker image build -t sleep-client:1.0.2 . && make docker.rmi.builder

docker.rmi.builder:
	docker image prune --filter label=stage=builder <<< y
