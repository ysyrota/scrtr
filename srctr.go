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
	Checked    *time.Time    `yaml:",omitempty"`
	Updated    *time.Time    `yaml:",omitempty"`
	Processing []replacement `yaml:",omitempty"`
}

const configFileName string = "srctr.yml"

func loadConfig(entries *[]listEntry) bool {
	entriesFile, err := os.Open(configFileName)
	if err == nil {
		decoder := yaml.NewDecoder(entriesFile)
		if err := decoder.Decode(&entries); err != nil {
			log.Fatalln("Error: ", err)
		}
		entriesFile.Close()
	} else if os.IsExist(err) {
		log.Println("Failed to open ", configFileName, ": ", err)
		return false
	}
	return true
}

func saveConfig(entries *[]listEntry) {
	ret, err := yaml.Marshal(entries)
	if err != nil {
		log.Fatalf("error: %v", err)
	}
	entriesFile, err := os.Create(configFileName)
	_, err = entriesFile.Write(ret)
	if err == nil {
		err = entriesFile.Sync()
	}
	if err == nil {
		err = entriesFile.Close()
	}
	if err != nil {
		log.Fatalf("Failed to write config: %v\n", err)
	}
}

func printList(entries *[]listEntry) {
	for _, e := range *entries {
		fmt.Printf("%v: checked %s, updated %s\n", e.Name, e.Checked, e.Updated)
	}
}

func main() {
	var entries []listEntry // = {
	//		{"google", "http://www.google.com/", time.Now(), time.Now(), []replacement{{"src", "dst"}}},
	//	}
	if !loadConfig(&entries) {
		os.Exit(1)
	}
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "list":
			printList(&entries)
			os.Exit(0)
		case "check":
		case "update":
			saveConfig(&entries)
			log.Fatalln("Not implemented")
		}
	}
	fmt.Println("Usage: srcrt [list|check <name>|update <name>]")
}
