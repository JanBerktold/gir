package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var CDCmd = &cobra.Command{
	Use:   "cd <owner/name>",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {

		data, err := LoadData()
		if err != nil {
			fmt.Printf("Failed to load data with error: %q\n", err.Error())
			os.Exit(-1)
		}

		if len(args) != 1 {
			fmt.Println("Invalid number of arguments")
			os.Exit(-1)
		}

		targetOwner, targetName := ParseRepo(args[0])
		if len(targetOwner) == 0 {
			fmt.Printf("Invalid repo name %s/%s\n", args[0])
			os.Exit(-1)
		}

		for _, repo := range data.Repositories {
			if repo.Owner.Name == targetOwner && repo.Name == targetName {
				data.CurrentRepo = &RepoIdentifier{
					Owner: targetOwner,
					Name:  targetName,
				}
				goto finished
			}

		}

		fmt.Printf("Repository %s/%s is not cached. You may want to run gir add %s/%s\n",
			targetOwner, targetName, targetOwner, targetName)
		return

	finished:
		if err := SaveData(data); err != nil {
			fmt.Printf("Failed to save data: %q\n", err.Error())
			os.Exit(-1)
		}
	},
}

func init() {
	RootCmd.AddCommand(CDCmd)
}
