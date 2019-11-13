package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

// list						list of tracking resources
// check [<name>|--all]
// update [<name>|--all]

type replacement struct {
	Search      string
	Replacement string
}
type listEntry struct {
	Name       string
	URL        string
	Checked    time.Time     `yaml:",omitempty"`
	Updated    time.Time     `yaml:",omitempty"`
	Processing []replacement `yaml:",omitempty"`
}

func save(entriesFileName string, entries []listEntry) {
	ret, err := yaml.Marshal(entries)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	entriesFile, err := os.Create(entriesFileName)
	_, err = entriesFile.Write(ret)
	if err != nil {
		log.Fatalf("Failed to write config: %v\n", err)
	}
	entriesFile.Close()
}

func (e listEntry) String() string {
	return fmt.Sprintf("%v: checked %s, updated %s", e.Name, e.Checked.Format(time.RFC3339), e.Updated)
}

func main() {
	var entries []listEntry // = {
	//		{"google", "http://www.google.com/", time.Now(), time.Now(), []replacement{{"src", "dst"}}},
	//	}
	entriesFileName := "entries.yml"
	entriesFile, err := os.Open(entriesFileName)
	if err == nil {
		decoder := yaml.NewDecoder(entriesFile)
		if err := decoder.Decode(&entries); err != nil {
			log.Fatalln("Error: ", err)
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
			save(entriesFileName, entries)
			log.Fatalln("Not implemented")
		}
	}
	fmt.Println("Usage: srcrt [list|check <name>|update <name>]")
}
