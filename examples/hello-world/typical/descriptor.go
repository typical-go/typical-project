package typical

import (
	"github.com/typical-go/typical-go/examples/hello-world/helloworld"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Descriptor of sample
var Descriptor = typcore.Descriptor{
	Name:    "hello-world",
	Version: "0.0.1",

	App: helloworld.New(), // the application

	BuildTool: typbuildtool.
		Create(
			typbuildtool.CreateModule(), // standard build module
		),
}
