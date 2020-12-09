package app

import (
	"github.com/go-openapi/runtime/middleware"
	apiHello "github.com/kozmod/idea-tests/open-api/client/internal/generated/restapi/operations/hello"
)

func (srv *Service) HelloWorldFullHandler(params apiHello.HelloWorldFullParams) middleware.Responder {
	return middleware.NotImplemented("operation hello HelloWorldFull has not yet been implemented")
}
