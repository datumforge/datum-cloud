package datumcloud

import (
	"github.com/schollz/progressbar/v3"
	"github.com/spf13/cobra"
)

// BarAdd is a wrapper around progressbar.Add() that handles error checking
func BarAdd(bar *progressbar.ProgressBar, num int) {
	err := bar.Add(num)
	cobra.CheckErr(err)
}
