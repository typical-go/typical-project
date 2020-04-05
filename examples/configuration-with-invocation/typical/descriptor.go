package typical

import (
	"github.com/typical-go/typical-go/examples/configuration-with-invocation/server"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typbuildtool"
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Descriptor of sample
var Descriptor = typcore.Descriptor{
	Name:    "configuration-with-invocation",
	Version: "1.0.0",

	App: typapp.EntryPoint(server.Main, "server").
		Configures(
			server.Configuration(), // Append configurer for the this project
		),

	BuildTool: typbuildtool.
		BuildSequences(
			typbuildtool.StandardBuild(),
		),
}
