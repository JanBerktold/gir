package main

import (
	"github.com/spf13/cobra"
)

var UpdateCmd = &cobra.Command{
	Use:   "update [<owner/name>]",
	Short: "Updates cached data for specified repository.
If not repository is given, all data is updated.
	",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	RootCmd.AddCommand(UpdateCmd)
}
