package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var ReposCmd = &cobra.Command{
	Use:   "repos",
	Short: "Gives an overview of all cached repositories",
	Run: func(cmd *cobra.Command, args []string) {

		data, err := LoadData()
		if err != nil {
			fmt.Printf("Failed to load data with error: %q\n", err.Error())
			os.Exit(-1)
		}

		for _, repo := range data.Repositories {
			fmt.Printf("%s/%s (%d Issues)\n", repo.Owner.Name, repo.Name, len(repo.Issues.Issues))
		}
	},
}

func init() {
	RootCmd.AddCommand(ReposCmd)
}
