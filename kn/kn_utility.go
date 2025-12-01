package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func utilityAction(c string) {
	fmt.Printf("Creating new %s utility...\n", c)

	utilityTemplate := `package utilities

func __FUNCNAME__() string {
	// TODO
	return ""
}
`

	file := strings.TrimSuffix(c, filepath.Ext(c))
	s1 := strings.Replace(utilityTemplate, "__FUNCNAME__", cases.Title(language.English).String(file), 1)

	if createNewFile("./internal/utilities/"+c+".go", s1) {
		fmt.Println("Utility created successfully!")
		//tidyAction()
	} else {
		fmt.Println("Failed to create utility.")
	}
}
