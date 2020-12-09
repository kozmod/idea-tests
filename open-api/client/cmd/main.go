package main

import (
	"log"

	"github.com/kozmod/idea-tests/open-api/client/internal/config"
	"github.com/kozmod/idea-tests/open-api/client/internal/generated/restapi"
	"github.com/kozmod/idea-tests/open-api/client/internal/generated/restapi/operations"

	apiHello "github.com/kozmod/idea-tests/open-api/client/internal/generated/restapi/operations/hello"

	"github.com/go-openapi/loads"

	"github.com/kozmod/idea-tests/open-api/client/internal/app"
)

func main() {
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	srv := app.New()
	api := operations.NewClientAPI(swaggerSpec)

	api.HelloHelloWorldHandler = apiHello.HelloWorldHandlerFunc(srv.HelloWorldHandler)
	api.HelloHelloWorldFullHandler = apiHello.HelloWorldFullHandlerFunc(srv.HelloWorldFullHandler)
	api.ServerShutdown = srv.OnShutdown
	server := restapi.NewServer(api)
	defer server.Shutdown()

	cfg, err := config.InitConfig("client")
	if err != nil {
		log.Fatalln(err)
	}

	server.ConfigureAPI()

	server.Port = cfg.HTTPBindPort
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
