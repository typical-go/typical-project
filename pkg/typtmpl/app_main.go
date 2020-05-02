package typtmpl

import (
	"io"
)

var _ Template = (*AppMain)(nil)

const appMain = `package main

// Autogenerated by Typical-Go. DO NOT EDIT.

import (
	"{{.DescPkg}}"
	"github.com/typical-go/typical-go/pkg/typcore"
)

func main() {
	typcore.LaunchApp(&typical.Descriptor)
}
`

// AppMain is writer to generate main.go for app
type AppMain struct {
	DescPkg string
}

// Execute app main template
func (t *AppMain) Execute(w io.Writer) (err error) {
	return Execute("appMain", appMain, t, w)
}
