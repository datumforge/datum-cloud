package main

import (
	datumcloud "github.com/datumforge/datum-cloud/cmd/cli/cmd"

	_ "github.com/datumforge/datum-cloud/cmd/cli/cmd/seed"
	_ "github.com/datumforge/datum-cloud/cmd/cli/cmd/workspace"
)

func main() {
	datumcloud.Execute()
}
