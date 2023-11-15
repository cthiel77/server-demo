package main

import (
	"github.com/cthiel77/server-demo/cmd"
	"github.com/cthiel77/server-demo/config"
)

// @title server-demo
// @version 1.0
// @description This is a demonstration application showing a working example of several functions and performance features.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email support@thiel-inet.de
// @license.name Apache 2.0
// @license.url https://opensource.org/license/bsd-3-clause/
// @BasePath /api
func main() {
	if flagsProcessed := config.Init(); !flagsProcessed {
		cmd.ExecuteCmd()
	}
}
