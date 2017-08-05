package main

import (
	"fmt"
)

func printHelp() {
	println("Whoops, looks like you forgot an argument!")
	println("Try:")
	println("  - help")
	println("  - name")
}

func startRailsServer() {
	println("One day this will start a Rails server.")
}

func projectName() {
	println(ProjDef.Name)
}

func printProjDef() {
	fmt.Printf("%v", ProjDef)
}
