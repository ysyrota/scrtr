package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"regexp"
	"time"
)

// list	list of tracking resources
// check [<name>|--all]
// update [<name>|--all]

type replacement struct {
	Search      regexp.Regexp `json:"source"`
	Replacement string        `json:"replacement"`
}
type listEntry struct {
	Name       string        `json:"name"`
	URL        string        `json:"url"`
	Checked    time.Time     `json:"checked"`
	Updated    time.Time     `json:"updated"`
	Processing []replacement `json:"replacement"`
}

func (e listEntry) String() string {
	return fmt.Sprintf("%v: checked %s, updated %s", e.Name, e.Checked.Format(time.RFC3339), e.Updated)
}

func main() {
	entries := []listEntry{}
	user, err := user.Current()
	if err != nil {
		fmt.Println("Failed to obtain current user: ", err)
		os.Exit(1)
	}
	entriesFileName := filepath.Join(user.HomeDir, "entries.json")
	entriesFile, err := os.Open(entriesFileName)
	if err == nil {
		fmt.Println("Reading entries from file")
		decoder := json.NewDecoder(entriesFile)
		if err := decoder.Decode(&entries); err != nil {
			fmt.Println("Error: ", err)
		}
		entriesFile.Close()
	} else if os.IsExist(err) {
		fmt.Println("Failed to open ", entriesFileName, ": ", err)
		os.Exit(1)
	}
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "list":
			fmt.Println(entries)
			os.Exit(0)
		case "check":
		case "update":
			fmt.Println("Not implemented")
			os.Exit(0)
		}
	}
	fmt.Println("Usage: srcmon [list|check <name>|update <name>]")
}
