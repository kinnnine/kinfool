package main

import (
	"fmt"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func routeAction(file string, method string) {

	checkArg(file, "route filename.")
	checkArg(method, "http method.")

	fmt.Printf("Creating new %s route...\n", file)

	routeTemplate := `package routes

import (
	"github.com/gin-gonic/gin"
)

func __FUNCNAME__(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "ok",
	})
}
`

	content := strings.Replace(routeTemplate, "__FUNCNAME__", cases.Title(language.English).String(file), 1)
	if createNewFile("./internal/routes/"+file+"."+method+".go", content) && createNewFile("./internal/controllers/"+file+".go", "") && createNewFile("./internal/services/"+file+".go", "") {
		fmt.Println("Route created successfully!")
		tidyAction()
	} else {
		fmt.Println("Failed to create route.")
	}
}
