package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Data struct {
	CurrentRepo  string
	Repositories []Repository
}

var file = os.Getenv("HOME") + "/.github_issues"

func LoadData() (Data, error) {
	f, err := os.Open(file)
	if err != nil {
		return Data{}, err
	}
	defer f.Close()

	var data Data
	return data, json.NewDecoder(f).Decode(&data)
}

func SaveData(data Data) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	return json.NewEncoder(f).Encode(data)
}

func ParseRepo(repo string) (owner, name string) {
	index := strings.Index(repo, "/")
	if index > -1 {
		return repo[0:index], repo[index+1:]
	}
	return "", ""
}

func main() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}
