package typical

import (
	"os"

	"github.com/typical-go/typical-go/pkg/buildkit"
	"github.com/typical-go/typical-go/pkg/typgo"
	"github.com/urfave/cli/v2"
)

func taskExamples(c *typgo.BuildCli) []*cli.Command {
	return []*cli.Command{
		{
			Name:    "test-example",
			Aliases: []string{"e"},
			Usage:   "Test all example",
			Action: func(cliCtx *cli.Context) (err error) {
				gotest := &buildkit.GoTest{
					Targets: []string{"./examples/..."},
				}

				cmd := gotest.Command()
				cmd.Stdout = os.Stdout
				cmd.Stderr = os.Stderr

				return cmd.Run(cliCtx.Context)
			},
		},
	}
}
