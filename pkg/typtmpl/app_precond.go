package typtmpl

import (
	"io"
)

var _ Template = (*AppPrecond)(nil)

const appPrecond = `typapp.Provide({{range $c := .Ctors}}
	&typapp.Constructor{Name: "{{$c.Name}}", Fn: {{$c.Def}}},{{end}}{{range $c := .CfgCtors}}
	&typapp.Constructor{
		Name: "{{$c.Name}}", 
		Fn: func() (cfg {{$c.SpecType}}, err error) {
			cfg = new({{$c.SpecType2}})
			if err = typcfg.Process("{{$c.Prefix}}", cfg); err != nil {
				return nil, err
			}
			return
		},
	},{{end}}
)
typapp.Destroy({{range $d := .Dtors}}
	&typapp.Destructor{Fn: {{$d.Def}}},{{end}}
)`

type (
	// AppPrecond to generate provide constructor
	AppPrecond struct {
		Ctors    []*Ctor
		CfgCtors []*CfgCtor
		Dtors    []*Dtor
	}

	// Ctor is constructor model
	Ctor struct {
		Name string
		Def  string
	}

	// Dtor is destructor model
	Dtor struct {
		Def string
	}

	// CfgCtor is config constructor model
	CfgCtor struct {
		Name      string
		Prefix    string
		SpecType  string
		SpecType2 string
	}
)

// Execute app precondition template
func (t *AppPrecond) Execute(w io.Writer) (err error) {
	return Execute("appPrecond", appPrecond, t, w)
}

// NotEmpty return true if not empty
func (t *AppPrecond) NotEmpty() bool {
	return len(t.Ctors) > 0 ||
		len(t.CfgCtors) > 0 ||
		len(t.Dtors) > 0
}
