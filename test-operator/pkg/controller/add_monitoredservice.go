package controller

import (
	"github.com/kozmod/idea-tests/test-operator/pkg/controller/monitoredservice"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, monitoredservice.Add)
}
