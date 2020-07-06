package typapp

import (
	"log"
	"os"

	"github.com/typical-go/typical-go/pkg/common"
	"github.com/typical-go/typical-go/pkg/typgo"
	"go.uber.org/dig"
)

type App struct {
	EntryPoint interface{}
}

func Start(entryPoint interface{}) {
	app := &App{
		EntryPoint: entryPoint,
	}

	if err := app.Run(); err != nil {
		log.Fatal(err.Error())
	}
}

// Run application
func (a *App) Run() error {
	if configFile := os.Getenv("CONFIG"); configFile != "" {
		if _, err := typgo.LoadConfig(configFile); err != nil {
			return err
		}
	}

	di := dig.New()
	for _, c := range _ctors {
		if err := di.Provide(c.Fn, dig.Name(c.Name)); err != nil {
			return err
		}
	}

	errs := common.GracefulRun(a.startFn(di), a.stopFn(di))
	return errs.Unwrap()
}

func (a *App) startFn(di *dig.Container) common.Fn {
	return func() (err error) {
		return di.Invoke(a.EntryPoint)
	}
}

func (a *App) stopFn(di *dig.Container) common.Fn {
	return func() (err error) {
		for _, dtor := range _dtors {
			if err = di.Invoke(dtor.Fn); err != nil {
				return
			}
		}
		return
	}
}
