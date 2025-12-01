package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func routeAction(c string) {
	fmt.Printf("Creating new %s route...\n", c)

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

	file := strings.TrimSuffix(c, filepath.Ext(c))
	s1 := strings.Replace(routeTemplate, "__FUNCNAME__", cases.Title(language.English).String(file), 1)
	method := strings.Split(c, ".")
	r1 := strings.Replace(c, "."+method[1], "", 1)

	if createNewFile("./internal/routes/"+c+".go", s1) && createNewFile("./internal/controllers/"+r1+".go", "") && createNewFile("./internal/services/"+r1+".go", "") {
		fmt.Println("Route created successfully!")
		tidyAction()
	} else {
		fmt.Println("Failed to create route.")
	}
}
