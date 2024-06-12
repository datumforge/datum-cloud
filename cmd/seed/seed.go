package datumseed

import (
	"github.com/spf13/cobra"

	datumcloud "github.com/datumforge/datum-cloud/cmd"
)

// seedCmd represents the base seed command when called without any subcommands
var seedCmd = &cobra.Command{
	Use:   "seed",
	Short: "the subcommands for creating demo data in datum",
}

func init() {
	datumcloud.RootCmd.AddCommand(seedCmd)
}
