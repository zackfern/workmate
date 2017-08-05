package main

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
