package main

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "gir",
	Short: "gir is a very simple, fast and flexible github offline issues reader.",
	Long: `gir is an offline github issues reader by github.com/JanBerktold
		Go to https://github.com/JanBerktold/gir for further information.`,
}
