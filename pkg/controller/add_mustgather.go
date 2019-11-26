package controller

import (
	"github.com/redhat-cop/must-gather-operator/pkg/controller/mustgather"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, mustgather.Add)
}
