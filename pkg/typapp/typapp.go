package typapp

import (
	log "github.com/sirupsen/logrus"
	"github.com/typical-go/typical-go/pkg/typcli"
	"github.com/typical-go/typical-go/pkg/typmod"

	"os"

	"github.com/typical-go/typical-go/pkg/typctx"
	"github.com/urfave/cli"
)

// Run the application
func Run(ctx *typctx.Context) {
	ctxCli := &typcli.ContextCli{
		Context: ctx,
	}
	app := cli.NewApp()
	app.Name = ctx.Name
	app.Usage = ""
	app.Description = ctx.Description
	app.Version = ctx.Version
	if actionable, ok := ctx.AppModule.(typmod.Actionable); ok {
		app.Action = ctxCli.Action(actionable.Action())
	}
	app.Before = typcli.LoadEnvFile
	if commander, ok := ctx.AppModule.(typcli.AppCommander); ok {
		app.Commands = commander.AppCommands(ctxCli)
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatal(err.Error())
	}
}
