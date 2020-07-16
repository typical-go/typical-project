package typtmpl

import (
	"io"
)

type (
	// ConfigAnnotated template
	ConfigAnnotated struct {
		Package  string
		Imports  []string
		CfgCtors []*CfgCtor
	}
	// CfgCtor is config constructor model
	CfgCtor struct {
		Name      string
		Prefix    string
		SpecType  string
		SpecType2 string
	}
)

var _ Template = (*ConfigAnnotated)(nil)

const configAnnotated = `package {{.Package}}

// Autogenerated by Typical-Go. DO NOT EDIT.

import ({{range $import := .Imports}}
	"{{$import}}"{{end}}
)

func init() { {{if .CfgCtors}}
	typapp.AppendCtor({{range $c := .CfgCtors}}
		&typapp.Constructor{
			Name: "{{$c.Name}}",
			Fn: func() (cfg {{$c.SpecType}}, err error) {
				cfg = new({{$c.SpecType2}})
				if err = typgo.ProcessConfig("{{$c.Prefix}}", cfg); err != nil {
					return nil, err
				}
				return
			},
		},{{end}}
	){{end}}
}`

// Execute app precondition template
func (t *ConfigAnnotated) Execute(w io.Writer) (err error) {
	return Parse("appPrecond", configAnnotated, t, w)
}
