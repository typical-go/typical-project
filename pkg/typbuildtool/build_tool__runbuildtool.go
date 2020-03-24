package typbuildtool

import (
	"os"

	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/urfave/cli/v2"
)

// RunBuildTool to run the build-tool
func (b *TypicalBuildTool) RunBuildTool(tc *typcore.Context) (err error) {
	c := &Context{
		Context:          tc,
		TypicalBuildTool: b,
	}

	app := cli.NewApp()
	app.Name = c.Name
	app.Usage = "Build-Tool"
	app.Description = c.Description

	app.Description = c.Description
	app.Before = func(cliCtx *cli.Context) (err error) {
		return b.Precondition(&BuildContext{
			Context: c,
			Cli:     cliCtx,
		})
	}
	app.Version = c.Version
	app.Commands = b.Commands(c)

	return app.Run(os.Args)
}
