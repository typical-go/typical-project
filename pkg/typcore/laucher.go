package typcore

import "log"

// LaunchApp the application
func LaunchApp(launcher AppLauncher) {
	if err := launcher.LaunchApp(); err != nil {
		log.Fatal(err.Error())
	}
}

// LaunchBuildTool the build tool
func LaunchBuildTool(launcher BuildToolLauncher) {
	if err := launcher.LaunchBuildTool(); err != nil {
		log.Fatal(err.Error())
	}
}
