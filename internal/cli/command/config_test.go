package command_test

import (
	"testing"

	"github.com/cthiel77/server-demo/internal/app"
	"github.com/cthiel77/server-demo/internal/cli/command"
	"github.com/spf13/cobra"
)

func TestCreatePrintConfigCmd(t *testing.T) {
	rootCmd := command.CreateRootCmd()

	// add print config cmd
	rootCmd.AddCommand(command.CreatePrintConfigCmd(
		func(cmd *cobra.Command, args []string) {
			app.PrintConfig()
		}))

}
