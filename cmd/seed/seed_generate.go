package datumseed

import (
	"github.com/spf13/cobra"

	datumcloud "github.com/datumforge/datum-cloud/cmd"
	"github.com/datumforge/datum-cloud/internal/seed"
)

var (
	defaultObjectCount = 10
	defaultInviteCount = 5
)

var seedGenerateCmd = &cobra.Command{
	Use:   "generate",
	Short: "generate random data for seeded environment with a single root organization",
	RunE: func(cmd *cobra.Command, args []string) error {
		return generate()
	},
}

func init() {
	seedCmd.AddCommand(seedGenerateCmd)

	seedGenerateCmd.Flags().StringP("directory", "d", "demodata", "directory to save generated data")
	seedGenerateCmd.Flags().Int("users", defaultObjectCount, "number of users to generate")
	seedGenerateCmd.Flags().Int("groups", defaultObjectCount, "approximate number of groups to generate")
	seedGenerateCmd.Flags().Int("invites", defaultInviteCount, "number of invites to generate")
}

func generate() error {
	config, err := seed.NewDefaultConfig()
	cobra.CheckErr(err)

	if datumcloud.Config.String("directory") != "" {
		config.Directory = datumcloud.Config.String("directory")
	}

	config.NumUsers = datumcloud.Config.Int("users")
	config.NumGroups = datumcloud.Config.Int("groups")
	config.NumInvites = datumcloud.Config.Int("invites")

	return config.GenerateData()
}
