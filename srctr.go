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

type source struct {
	URL        string
	Checked    *time.Time    `yaml:",omitempty"`
	Updated    *time.Time    `yaml:",omitempty"`
	Processing []replacement `yaml:",omitempty"`
}

type sources map[string]source

const configFileName string = "srctr.yml"

func loadConfig(cfg *sources) bool {
	entriesFile, err := os.Open(configFileName)
	if err == nil {
		decoder := yaml.NewDecoder(entriesFile)
		if err := decoder.Decode(&cfg); err != nil {
			log.Fatalln("Error: ", err)
		}
		entriesFile.Close()
	} else if os.IsExist(err) {
		log.Println("Failed to open ", configFileName, ": ", err)
		return false
	}
	return true
}

func saveConfig(cfg sources) {
	ret, err := yaml.Marshal(cfg)
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

func printList(src sources) {
	for name, e := range src {
		fmt.Printf("%v: checked %s, updated %s\n", name, e.Checked, e.Updated)
	}
}

func main() {
	var cfg sources

	if !loadConfig(&cfg) {
		os.Exit(1)
	}
	if len(os.Args) > 1 {
		switch os.Args[1] {
		case "list":
			printList(cfg)
			os.Exit(0)
		case "check":
		case "update":
			saveConfig(cfg)
			log.Fatalln("Not implemented")
		}
	}

	fmt.Println("Usage: srcrt [list|check <name>|update <name>]")
}
