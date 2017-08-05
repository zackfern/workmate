package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// FileName defines the name of the JSON file we're hoping to find.
const FileName = "workmate.json"

// WorkMate mimics the JSON structure we're parsing and sets accessors/types.
type WorkMate struct {
	Name        string `json:"name"`
	Port        int    `json:"port"`
	PullOnEntry bool   `json:"pull_on_entry"`
}

var ProjDef WorkMate

func main() {
	// Read the file. This will exit if it's not found.
	file := findFile()
	// Parse the file...
	parseFile(file)

	// If they didn't pass an argument, run the hooks and print help.
	if len(os.Args) == 1 {
		if anyHooks() {
			runHooks()
		} else {
			printHelp()
		}
		os.Exit(0)
	}

	// Figure out which command somebody is trying to run...
	command := os.Args[1]

	switch command {
	case "help":
		printHelp()
	case "start":
		startRailsServer()
	case "name":
		projectName()
	case "print":
		printProjDef()
	case "branch":
		printGitBranch()
	}
}

func findFile() []byte {
	dat, err := ioutil.ReadFile(FileName)

	// If there's an error, exit the program.
	if err != nil {
		fmt.Printf("Couldn't find %s defintion file!", FileName)
		os.Exit(1)
	}

	return dat
}

func parseFile(file []byte) {
	json.Unmarshal(file, &ProjDef)
}
