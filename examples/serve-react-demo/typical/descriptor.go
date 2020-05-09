package typical

import (
	"github.com/typical-go/typical-go/examples/serve-react-demo/server"
	"github.com/typical-go/typical-go/pkg/typapp"
	"github.com/typical-go/typical-go/pkg/typbuild"
	"github.com/typical-go/typical-go/pkg/typcore"
)

// Descriptor of sample
var Descriptor = typcore.Descriptor{
	Name:    "server-echo-react",
	Version: "1.0.0",

	App: &typapp.App{
		EntryPoint: server.Main,
	},

	BuildTool: &typbuild.BuildTool{
		BuildSequences: []interface{}{
			&ReactDemoModule{source: "react-demo"},
			typbuild.StandardBuild(),
		},
	},
}
