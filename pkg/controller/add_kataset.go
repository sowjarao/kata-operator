package controller

import (
	"mydomain.com/mygroup2/kata-operator/pkg/controller/kataset"
)

func init() {
	// AddToManagerFuncs is a list of functions to create controllers and add them to a manager.
	AddToManagerFuncs = append(AddToManagerFuncs, kataset.Add)
}
