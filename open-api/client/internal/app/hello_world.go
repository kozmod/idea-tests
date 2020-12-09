package app

import (
	"github.com/go-openapi/runtime/middleware"
	apiHello "github.com/kozmod/idea-tests/open-api/client/internal/generated/restapi/operations/hello"
)

func (srv *Service) HelloWorldHandler(params apiHello.HelloWorldParams) middleware.Responder {
	return middleware.NotImplemented("operation hello HelloWorld has not yet been implemented")
}
