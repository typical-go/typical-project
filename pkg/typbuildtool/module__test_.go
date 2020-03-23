package typbuildtool

import (
	"fmt"

	"github.com/typical-go/typical-go/pkg/buildkit"
)

// Test the project
func (b *Module) Test(c *BuildContext) (err error) {
	var targets []string
	for _, source := range c.ProjectSources {
		targets = append(targets, fmt.Sprintf("./%s/...", source))
	}

	gotest := buildkit.NewGoTest(targets...).
		WithCoverProfile(b.coverProfile).
		WithRace(true).
		WithStdout(b.stdout).
		WithStderr(b.stderr)

	return gotest.Execute(c.Cli.Context)
}
