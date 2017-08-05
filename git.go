package main

import "os/exec"

func gitPull() {
	println("Pulling Git...")
	gitPull := exec.Command("git", "pull")

	out, runErr := gitPull.CombinedOutput()
	hErr(runErr)
	println("\t" + string(out))
}
