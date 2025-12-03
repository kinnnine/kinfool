package main

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func utilityAction(file string) {
	checkArg(file, "utility filename.")

	fmt.Printf("Creating new %s utility...\n", file)

	utilityTemplate := `package utilities

func __FUNCNAME__() string {
	// TODO
	return ""
}
`

	content := strings.Replace(utilityTemplate, "__FUNCNAME__", cases.Title(language.English).String(file), 1)
	if createNewFile("./internal/utilities/"+file+".go", content) {
		fmt.Println("Utility created successfully!")
		//tidyAction()
	} else {
		fmt.Println("Failed to create utility.")
	}
}
