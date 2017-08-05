package main

import (
	"os/exec"
)

func anyHooks() bool {
	if ProjDef.PullOnEntry {
		return true
	} else {
		return false
	}
}

func runHooks() {
	if ProjDef.PullOnEntry {
		gitPull()
	}
}

func gitBranchName() []byte {
	gitName := exec.Command("git symbolic-ref --short -q HEAD")
	data, _ := gitName.Output()
	return data
}

func gitStatus() []byte {
	gitClean := exec.Command("git status --untracked-files=no --porcelain")
	data, _ := gitClean.Output()
	return data
}

func gitPull() {
	clean := gitStatus()
	if clean == nil {
		println("Clean, should be able to pull.")
	} else {
		println("Can't pull, unclean.")
	}

	gitPull := exec.Command("git pull")
	out, _ := gitPull.Output()
	println(out)
}
