package main

import (
	datumcloud "github.com/datumforge/datum-cloud/cmd"

	_ "github.com/datumforge/datum-cloud/cmd/seed"
)

func main() {
	datumcloud.Execute()
}
