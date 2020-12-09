roma:
	swagger generate server -f swagger/swagger.yml --default-scheme=http -m=app/domain/model --server-package=app/interface/rest/api

generate:
	swagger generate server -f swagger/swagger.yml --name example

tmpl:
	@rm -rf ./cmd
	@rm -rf ./internal
	@mkdir -p cmd
	@mkdir -p internal/generated
	@mkdir -p internal/app
	@mkdir -p internal/config
	@swagger generate server \
    		-f ./swagger/swagger.yml \
    		-t ./internal/generated -C ./swagger-templates/default-server.yml \
    		--template-dir ./swagger-templates/templates \
    		--name client