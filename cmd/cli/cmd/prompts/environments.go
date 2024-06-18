package prompts

import (
	"strings"

	"github.com/manifoldco/promptui"
)

type env struct {
	Name   string
	Values []string
}

var environments = []env{
	{
		Name:   "Production & Testing",
		Values: []string{"production", "testing"},
	},
	{
		Name:   "Production Only",
		Values: []string{"production"},
	},
	{
		Name:   "Testing Only",
		Values: []string{"testing"},
	},
}

var selectTemplates = &promptui.SelectTemplates{
	Label:    "{{ .Name }} ",
	Active:   "\U0001F449 {{ .Name | cyan }}",
	Inactive: "  {{ .Name | cyan }}",
	Selected: "\U0001F449 {{ .Name | green | cyan }}",
	Details: `
{{ "Name:" | faint }}	{{ .Name }}
{{ "Values:" | faint }}	{{ .Values }}`,
}

var searcher = func(input string, index int) bool {
	environments := environments[index]
	name := strings.ReplaceAll(strings.ToLower(environments.Name), " ", "")
	input = strings.ReplaceAll(strings.ToLower(input), " ", "")

	return strings.Contains(name, input)
}

func Environments() ([]string, error) {
	prompt := promptui.Select{
		Label:     "Environments:",
		Templates: selectTemplates,
		Items:     environments,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	return environments[i].Values, err
}
