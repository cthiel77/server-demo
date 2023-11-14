package command

import (
	"github.com/cthiel77/server-demo/internal/cli/branding"
	"github.com/cthiel77/server-demo/internal/meta"
	"github.com/spf13/cobra"
)

// CreateRootCmd creates a root command node to attach cli commands
func CreateRootCmd() cobra.Command {
	return cobra.Command{
		Use:     meta.GetAppName(),
		Short:   meta.GetAppDescription(),
		Long:    branding.GetIntro(),
		Version: meta.GetVersionID(),
	}
}
