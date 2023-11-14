package command

import "github.com/spf13/cobra"

// CreatePrintConfigCmd creates a print config command
func CreatePrintConfigCmd(fn func(cmd *cobra.Command, args []string)) *cobra.Command {
	return &cobra.Command{
		Use:   "config",
		Short: "print the config and exit",
		Run:   fn,
	}
}
