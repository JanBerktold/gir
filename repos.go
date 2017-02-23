package main

import (
	"github.com/spf13/cobra"
)

var ReposCmd = &cobra.Command{
	Use:   "repos",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	RootCmd.AddCommand(ReposCmd)
}
