// Code generated by MockIO. DO NOT EDIT.
// Source: sample/mockio/interface.go

// Package mockio_foo is a generated GoMock package.
package mockio_foo

import (
	foo "github.com/golang/mock/sample/mockio"
)

type Get struct {
	Foo  int
	Bar  string
	Ret0 foo.Model
	Ret1 error
}

type List struct {
	Foos []int
	Bars []string
	Ret0 []foo.Model
	Ret1 error
}
