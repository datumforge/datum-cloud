package datumworkspace

import (
	"context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"

	datumcloud "github.com/datumforge/datum-cloud/cmd/cli/cmd"
	"github.com/datumforge/datum-cloud/cmd/cli/cmd/prompts"
	"github.com/datumforge/datum-cloud/internal/v1/models"
)

var workspaceCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new datum org",
	RunE: func(cmd *cobra.Command, args []string) error {
		return createWorkspace(cmd.Context())
	},
}

func init() {
	workspaceCmd.AddCommand(workspaceCreateCmd)

	workspaceCreateCmd.Flags().StringP("name", "n", "", "name of the workspace")
	workspaceCreateCmd.Flags().StringP("description", "d", "", "description of the workspace")
	workspaceCreateCmd.Flags().StringSlice("domains", []string{}, "domains associated with the workspace")
	workspaceCreateCmd.Flags().BoolP("interactive", "i", true, "interactive prompt, set to false to disable")
}

func createWorkspace(ctx context.Context) error {
	c, err := datumcloud.SetupClient(ctx, datumcloud.Config.String("host"))
	cobra.CheckErr(err)

	// check if interactive flag is set, if it is, and some params are not set, prompt user for input
	interactive := datumcloud.Config.Bool("interactive")

	name := datumcloud.Config.String("name")
	if name == "" && interactive {
		name, err = prompts.Name()
		cobra.CheckErr(err)
	}

	description := datumcloud.Config.String("description")
	if description == "" && interactive {
		description, err = prompts.Description()
		cobra.CheckErr(err)
	}

	domains := datumcloud.Config.Strings("domains")
	if len(domains) == 0 && interactive {
		domainString, err := prompts.Domains()
		cobra.CheckErr(err)

		if domainString != "" {
			domains = strings.Split(domainString, ",")
		}
	}

	environments := datumcloud.Config.Strings("environments")
	if len(environments) == 0 && interactive {
		environments, err = prompts.Environments()
		cobra.CheckErr(err)

		fmt.Println("Environments: ", environments)
	}

	input := models.WorkspaceRequest{
		Name:         name,
		Description:  description,
		Domains:      domains,
		Environments: environments,
	}

	// add an empty line
	fmt.Println()

	// create progress bar
	bar := progressbar.NewOptions(100, //nolint:mnd
		progressbar.OptionEnableColorCodes(true),
		progressbar.OptionShowBytes(false),
		progressbar.OptionSetWidth(15), //nolint:mnd
		progressbar.OptionShowElapsedTimeOnFinish(),
		progressbar.OptionSetDescription("[light_green]>[reset] creating workspaces..."),
		progressbar.OptionSetTheme(progressbar.Theme{
			Saucer:        "[light_green]=[reset]",
			SaucerHead:    "[light_green]>[reset]",
			SaucerPadding: " ",
			BarStart:      "[",
			BarEnd:        "]",
		}),
	)
	defer bar.Exit() //nolint:errcheck

	var (
		wg sync.WaitGroup
		ws *models.WorkspaceReply
	)

	// create a wait channel
	waitCh := make(chan struct{})

	// add a wait group
	wg.Add(1)

	go func() {
		go func() {
			defer wg.Done()

			ws, err = c.WorkspaceCreate(ctx, &input)
			cobra.CheckErr(err)
		}()

		wg.Wait()
		close(waitCh)
	}()

	// Block until the wait group is done
	wait(waitCh, bar)

	fmt.Println("\nID: ", ws.ID)
	fmt.Println("New Workspace Created: ", ws.Name)

	if ws.Description != "" {
		fmt.Println("Description: ", ws.Description)
	}

	if len(ws.Domains) > 0 {
		fmt.Println("Domains: ", strings.Join(ws.Domains, ","))
	}

	for _, env := range ws.Environments {
		// add an empty line
		fmt.Println()

		fmt.Println("--> Environment: ", env.Name)

		for _, bucket := range env.Buckets {
			fmt.Println("-----> Bucket: ", bucket.Name)

			for _, relation := range bucket.Relations {
				fmt.Println("--------> Relation: ", relation.Name)
			}
		}

		// add an empty line
		fmt.Println()
	}

	return nil
}

// wait will wait for the wait group to finish and update the progress bar
func wait(waitCh chan struct{}, bar *progressbar.ProgressBar) {
	for {
		select {
		case <-waitCh:
			err := bar.Finish()
			cobra.CheckErr(err)

			return
		case <-time.After(100 * time.Millisecond): //nolint:mnd
			// update the progress bar while waiting
			datumcloud.BarAdd(bar, 1) //nolint:mnd
		}
	}
}
