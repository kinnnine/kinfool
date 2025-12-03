package main

import (
	"fmt"
	"os"
	"strings"
)

func initializeAction(name string) {

	checkArg(name, "project name.")

	urlSplit := strings.Split(name, "/")
	projectName := urlSplit[len(urlSplit)-1]

	fmt.Printf("Initializing empty kinfool project for %s...\n", name)

	internal_structure := [...]string{
		"controllers", "middlewares", "routes", "services", "utilities",
	}

	for _, value := range internal_structure {
		if !createNewFolder(projectName + "/internal/" + value) {
			os.Exit(1)
		}
	}

	os.Chdir("./" + projectName)

	createNewFile("./kinfool.go", "")
	tidyAction()

	runCmd("go", "mod", "init", name)
	runCmd("go", "mod", "tidy")
}
