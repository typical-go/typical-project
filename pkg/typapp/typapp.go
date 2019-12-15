package typapp

import (
	"os"

	log "github.com/sirupsen/logrus"

	"github.com/typical-go/typical-go/pkg/typcore"
	"github.com/typical-go/typical-go/pkg/utility/envfile"
	"github.com/urfave/cli/v2"
)

// Run the application
func Run(ctx *typcore.Context) {
	if err := ctx.Validate(); err != nil {
		log.Fatal(err.Error())
	}
	appCli := typcore.NewCli(ctx, ctx.AppModule)
	app := cli.NewApp()
	app.Name = ctx.Name
	app.Usage = ""
	app.Description = ctx.Description
	app.Version = ctx.Version
	if actionable, ok := ctx.AppModule.(typcore.Actionable); ok {
		app.Action = appCli.PreparedAction(actionable.Action())
	}
	app.Before = func(ctx *cli.Context) error {
		return envfile.Load()
	}
	if commander, ok := ctx.AppModule.(typcore.AppCommander); ok {
		app.Commands = commander.AppCommands(appCli)
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}
