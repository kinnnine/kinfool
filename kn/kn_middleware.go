package main

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func middlewareAction(file string) {
	checkArg(file, "middleware filename.")

	fmt.Printf("Creating new %s middleware...\n", file)

	middlewareTemplate := `package middlewares

import (
	"github.com/gin-gonic/gin"
)

func __FUNCNAME__() gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO
	}
}
`

	content := strings.Replace(middlewareTemplate, "__FUNCNAME__", cases.Title(language.English).String(file), 1)
	if createNewFile("./internal/middlewares/"+file+".go", content) {
		fmt.Println("Middleware created successfully!")
		tidyAction()
	} else {
		fmt.Println("Failed to create middleware.")
	}
}
