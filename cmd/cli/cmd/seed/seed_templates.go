package datumseed

import (
	"context"
	"fmt"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"

	datumcloud "github.com/datumforge/datum-cloud/cmd/cli/cmd"
	"github.com/datumforge/datum-cloud/internal/seed"
)

var seedTemplateCmd = &cobra.Command{
	Use:   "templates",
	Short: "add templates to an existing seeded environment",
	RunE: func(cmd *cobra.Command, args []string) error {
		return initTemplateData(cmd.Context())
	},
}

func init() {
	seedCmd.AddCommand(seedTemplateCmd)
}

func initTemplateData(ctx context.Context) error {
	c, err := newSeedClient()
	cobra.CheckErr(err)

	bar := progressbar.NewOptions(100, //nolint:mnd
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(false),
		progressbar.OptionSetWidth(15), //nolint:mnd
		progressbar.OptionShowElapsedTimeOnFinish(),
		progressbar.OptionSetDescription("[light_green]>[reset] creating seeded templates..."),
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

	bar.Describe("[light_green]>[reset] creating templates...")
	datumcloud.BarAdd(bar, 10) //nolint:mnd

	err = c.LoadTemplates(ctx)
	cobra.CheckErr(err)

	bar.Describe("[light_green]>[reset] seeded environment created")
	err = bar.Finish()
	cobra.CheckErr(err)

	return getAllTemplateData(ctx, c)
}

// getAllTemplateData gets all the data from the seeded environment in a table format
func getAllTemplateData(ctx context.Context, c *seed.Client) error {
	templates, err := c.GetAllTemplates(ctx)
	cobra.CheckErr(err)

	header := table.Row{"ID", "Name"}
	rows := []table.Row{}

	for _, template := range templates.Templates.Edges {
		rows = append(rows, []interface{}{template.Node.ID, template.Node.Name})
	}

	// add empty row for spacing
	fmt.Println()

	createTableOutput("Templates", header, rows)

	return nil
}
