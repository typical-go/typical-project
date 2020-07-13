package main

import (
	"log"

	"github.com/typical-go/typical-go/pkg/execkit"
	"github.com/typical-go/typical-go/pkg/typgo"
)

var (
	descriptor = typgo.Descriptor{
		Name:    "typical-go",
		Version: "0.9.57",
		Layouts: []string{"internal", "pkg"},

		Commands: typgo.Commands{
			&typgo.CompileCmd{
				Action: &typgo.StdCompile{MainPackage: "."},
			},
			&typgo.RunCmd{
				Action: &typgo.StdRun{},
			},
			&typgo.TestCmd{
				Action: &typgo.StdTest{},
			},
			&typgo.CleanCmd{
				Action: &typgo.StdClean{},
			},
			&typgo.ReleaseCmd{
				Releaser: &typgo.Github{Owner: "typical-go", Repo: "typical-go"},
			},

			&typgo.Command{
				Name:    "examples",
				Aliases: []string{"e"},
				Usage:   "Test all example",
				Action:  typgo.NewAction(testExamples),
			},
		},
	}
)

func testExamples(c *typgo.Context) error {
	return c.Execute(&execkit.GoTest{
		Targets: []string{"./examples/..."},
	})
}

func main() {
	if err := typgo.Run(&descriptor); err != nil {
		log.Fatal(err.Error())
	}
}
