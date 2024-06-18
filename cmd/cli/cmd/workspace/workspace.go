package datumworkspace

import (
	"github.com/spf13/cobra"

	datumcloud "github.com/datumforge/datum-cloud/cmd/cli/cmd"
)

// workspaceCmd represents the base workspace command when called without any subcommands
var workspaceCmd = &cobra.Command{
	Use:   "workspace",
	Short: "the subcommands for working with the datum workspace",
}

func init() {
	datumcloud.RootCmd.AddCommand(workspaceCmd)
}
