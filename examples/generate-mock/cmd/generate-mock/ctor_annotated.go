package main

// Autogenerated by Typical-Go. DO NOT EDIT.

import (
	"github.com/typical-go/typical-go/examples/generate-mock/internal/helloworld"
	"github.com/typical-go/typical-go/pkg/typapp"
)

func init() {
	typapp.AppendCtor(
		&typapp.Constructor{Name: "", Fn: helloworld.GetWriter},
		&typapp.Constructor{Name: "", Fn: helloworld.NewGreeter},
	)
}
