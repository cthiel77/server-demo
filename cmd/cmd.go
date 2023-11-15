package cmd

import (
	"fmt"
	"os"

	"github.com/cthiel77/server-demo/internal/cli/command"
)

var (
	rootCmd = command.CreateRootCmd()
)

// ExecuteCmd Inits defined cmds. This is called by main.main().
// It only needs to happen once to the rootCmd.
func ExecuteCmd() {
	// Execute adds all child commands to the root command and sets flags appropriately.
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
