package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var ListCmd = &cobra.Command{
	Use:   "list [<owner/name>]",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {

		data, err := LoadData()
		if err != nil {
			fmt.Printf("Failed to load data with error: %q\n", err.Error())
			os.Exit(-1)
		}

		var targetOwner, targetName string
		if len(args) == 1 {
			targetOwner, targetName = ParseRepo(args[0])
			if len(targetOwner) == 0 {
				fmt.Printf("Invalid repo name %q\n", args[0])
				os.Exit(-1)
			}
		} else if data.CurrentRepo != nil {
			targetOwner, targetName = data.CurrentRepo.Owner, data.CurrentRepo.Name
			if len(targetOwner) == 0 {
				fmt.Printf("Invalid repo name %q\n", args[0])
				os.Exit(-1)
			}
		} else {
			fmt.Println("No repository specified. You either have to specify a repo as an argument or enter a scope.\nSee gir list --help and gir cd --help for details")
			os.Exit(-1)
		}

		for _, repo := range data.Repositories {
			if repo.Owner.Name == targetOwner && repo.Name == targetName {
				for _, issue := range repo.Issues.Issues {
					fmt.Printf("%d/%s: %s (%s)\n", issue.Number, issue.State, issue.Title, issue.Author.Name)
				}
				return
			}
		}

	},
}

func init() {
	RootCmd.AddCommand(ListCmd)
}
