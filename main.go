package main

import (
	"github.com/cthiel77/server-demo/cmd"
	"github.com/cthiel77/server-demo/config"
)

func main() {
	if flagsProcessed := config.Init(); !flagsProcessed {
		cmd.ExecuteCmd()
	}
}
