package prompts

import (
	"github.com/manifoldco/promptui"

	datumcloud "github.com/datumforge/datum-cloud/cmd/cli/cmd"
)

func Name() (string, error) {
	validate := func(input string) error {
		if len(input) == 0 {
			return datumcloud.NewRequiredFieldMissingError("name")
		}

		return nil
	}

	prompt := promptui.Prompt{
		Label:     "Name:",
		Templates: templates,
		Validate:  validate,
	}

	return prompt.Run()
}
