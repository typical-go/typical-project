package typical

import (
	"github.com/typical-go/typical-go/examples/simple-additional-task/helloworld"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Descriptor of sample
var Descriptor = typcore.Descriptor{
	Name:    "simple-additional-task",
	Version: "0.0.1",

	App: helloworld.New(),

	BuildTool: typbuildtool.
		Create(
			typbuildtool.CreateModule(),
		).
		WithCommanders(
			typbuildtool.CreateCommander(taskPrintContext),
		),
}
