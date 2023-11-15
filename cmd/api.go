// Package cmd bus listening commands
package cmd

import (
	"github.com/spf13/cobra"

	"github.com/cthiel77/server-demo/config"
	"github.com/cthiel77/server-demo/server"
)

var runAPIServerCmd = &cobra.Command{
	Use:   config.CmdNameWebServer,
	Short: "Serves the API",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		server.RunAPIServer(args)
	},
}

func init() {
	rootCmd.AddCommand(runAPIServerCmd)
}
