package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func tidyAction() {
	checkKinfool()
	fmt.Println("Tidying kinfool.go...")

	// kinfool.go contents
	kinfoolTemplate := `package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	__MIDDLEWARES_IMPORT__
	__ROUTERS_IMPORT__
)

func main() {
	// Load environment variables from .env file (if exists)
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	listen := os.Getenv("LISTEN")
	port := os.Getenv("PORT")
	trustedProxies := os.Getenv("TRUSTED_PROXIES")

	// Create a Gin router
	r := gin.New()

	// Set trusted proxies from environment variable
	r.SetTrustedProxies([]string{"" + trustedProxies})

	// Apply middlewares
	__MIDDLEWARES__
	// Routers
	{

		__ROUTERS__
	}

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	r.Run(listen + ":" + port)
}
`

	s1 := ""

	// internal/middlewares
	if fileExists("./internal/middlewares") {
		files := getDirectoryContents("./internal/middlewares")

		if len(files) > 0 {
			s1 = strings.Replace(kinfoolTemplate, "__MIDDLEWARES_IMPORT__", `"github.com/kinnnine/kinfool/internal/middlewares"`, 1)
			var sb strings.Builder

			for _, value := range files {
				file := strings.TrimSuffix(value, filepath.Ext(value))
				sb.WriteString(fmt.Sprintf("r.Use(middlewares.%s())\n\t", cases.Title(language.English).String(file)))
			}
			s1 = strings.Replace(s1, "__MIDDLEWARES__", sb.String(), 1)
		} else {
			s1 = strings.Replace(kinfoolTemplate, "__MIDDLEWARES_IMPORT__", "", 1)
			s1 = strings.Replace(s1, "__MIDDLEWARES__", "", 1)
		}
	}

	// internal/routes
	if fileExists("./internal/routes") {
		files := getDirectoryContents("./internal/routes")

		if len(files) > 0 {
			s1 = strings.Replace(s1, "__ROUTERS_IMPORT__", `"github.com/kinnnine/kinfool/internal/routes"`, 1)
			var sb strings.Builder
			sb.WriteString("api := r.Group(\"/api\")\n\t\t")

			for _, value := range files {
				file1 := strings.TrimSuffix(value, filepath.Ext(value))
				file := strings.TrimSuffix(file1, filepath.Ext(file1))
				method := strings.TrimPrefix(filepath.Ext(file1), ".")
				sb.WriteString(fmt.Sprintf("api.%s(\"%s\", routes.%s)\n\t\t", strings.ToUpper(method), file, cases.Title(language.English).String(file)))
			}
			s1 = strings.Replace(s1, "__ROUTERS__", sb.String(), 1)
		} else {
			s1 = strings.Replace(s1, "__ROUTERS_IMPORT__", "", 1)
			s1 = strings.Replace(s1, "__ROUTERS__", "", 1)
		}
	}

	if updateKinfool(s1) {
		fmt.Println("Tidying successfully.")
	} else {
		fmt.Println("Tidying failed.")
	}
}
