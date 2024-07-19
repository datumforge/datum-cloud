package datumseed

import (
	"context"
	"fmt"
	"time"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"

	datumcloud "github.com/datumforge/datum-cloud/cmd/cli/cmd"
	"github.com/datumforge/datum-cloud/internal/seed"
	"github.com/datumforge/datum/pkg/datumclient"
)

var seedOrgMembersCmd = &cobra.Command{
	Use:   "org-members",
	Short: "add users to an existing seeded organization",
	RunE: func(cmd *cobra.Command, args []string) error {
		return initOrgMemberData(cmd.Context())
	},
}

func init() {
	seedCmd.AddCommand(seedOrgMembersCmd)

	seedOrgMembersCmd.Flags().StringP("organization-id", "o", "", "organization ID to add users to")
	seedOrgMembersCmd.Flags().Int("users", defaultObjectCount, "number of users to generate")
	seedOrgMembersCmd.Flags().StringP("patid", "t", "", "personal access token ID to authorize the organization")
}

func initOrgMemberData(ctx context.Context) error {
	patID := datumcloud.Config.String("patid")
	if patID == "" {
		cobra.CheckErr("PAT ID not provided")
	}

	orgID := datumcloud.Config.String("organization-id")
	if orgID == "" {
		cobra.CheckErr("Organization ID not provided")
	}

	c, err := newSeedClient()
	cobra.CheckErr(err)

	// generate users in csv
	config, err := seed.NewDefaultConfig()
	cobra.CheckErr(err)

	config.NumUsers = datumcloud.Config.Int("users")

	config.GenerateUserData()

	bar := progressbar.NewOptions(100, //nolint:mnd
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(false),
		progressbar.OptionSetWidth(15), //nolint:mnd
		progressbar.OptionShowElapsedTimeOnFinish(),
		progressbar.OptionSetDescription("[light_green]>[reset] creating seeded org members..."),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[light_green]=[reset]",
			SaucerHead:    "[light_green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
	)
	defer bar.Exit() //nolint:errcheck

	datumcloud.BarAdd(bar, 10) //nolint:mnd

	bar.Describe("[light_green]>[reset] registering users...")
	datumcloud.BarAdd(bar, 10) //nolint:mnd

	userIDs, err := c.RegisterUsers(ctx)
	cobra.CheckErr(err)

	err = c.AuthorizeOrganizationOnPAT(ctx, orgID, patID)
	cobra.CheckErr(err)

	// wait for the cache to update (1s)
	// otherwise the user will not be able to see the org
	time.Sleep(2 * time.Second) //nolint:mnd
	datumcloud.BarAdd(bar, 10)  //nolint:mnd

	// create API Token for the root org and authorize as that token
	err = c.GenerateSeedAPIToken(ctx, orgID)
	cobra.CheckErr(err)

	bar.Describe("[light_green]>[reset] creating org members...")
	datumcloud.BarAdd(bar, 10) //nolint:mnd

	err = c.LoadOrgMembers(ctx, userIDs)
	cobra.CheckErr(err)

	bar.Describe("[light_green]>[reset] seeded environment created")
	err = bar.Finish()
	cobra.CheckErr(err)

	return getAllOrgMemberData(ctx, c, orgID)
}

// getAllOrgMemberData gets all the data from the seeded environment in a table format
func getAllOrgMemberData(ctx context.Context, c *seed.Client, orgID string) error {
	members, err := c.GetOrgMembersByOrgID(ctx, &datumclient.OrgMembershipWhereInput{
		OrganizationID: &orgID,
	})
	cobra.CheckErr(err)

	header := table.Row{"ID", "Email", "Role"}
	rows := []table.Row{}

	for _, om := range members.OrgMemberships.Edges {
		rows = append(rows, []interface{}{om.Node.ID, om.Node.User.Email, om.Node.Role})
	}

	// add empty row for spacing
	fmt.Println()

	createTableOutput("OrgMembers", header, rows)

	return nil
}
