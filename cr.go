package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var CRCmd = &cobra.Command{
	Use:   "cr <owner/name>",
	Short: "Selects an active repository which gir list and gir show work against",
	Long: "gir cr <owner/name> allows you to set a default repository for gir list and gir show.\n" +
		"Example:\n" +
		"\t> gir repos\n" +
		"\t\tgolang/go (1600 issues)\n" +
		"\t> gir list\n" +
		"\t\tNo repository specified. You either have to specify a repo as an argument or enter a scope.\n" +
		"\t> gir cr golang/go\n" +
		"\t> gir list\n" +
		"\t\t1/CLOSED: Make Golang awesome (bradfitz)\n" +
		"\t\t2/CLOSED: Think about generics (robpike)\n" +
		"\t\t3/OPEN: Gir is a hip github issues reader (JanBerktold)\n" +
		"\t> gir show 3\n" +
		"\t\t> issues info here\n",
	Run: func(cmd *cobra.Command, args []string) {

		data, err := LoadData()
		if err != nil {
			fmt.Printf("Failed to load data with error: %q\n", err.Error())
			os.Exit(-1)
		}

		var targetOwner string
		var targetName string
		if len(args) > 0 {
			targetOwner, targetName = ParseRepo(args[0])
		}
		if len(targetOwner) == 0 {
			data.CurrentRepo = nil
			goto finished
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
	RootCmd.AddCommand(CRCmd)
}
