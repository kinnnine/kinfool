package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {

	cli.AppHelpTemplate = fmt.Sprintf(`%s
MANUAL: 
	 https://github.com/kinnnine/kinfool/wiki

`, cli.AppHelpTemplate)

	app := cli.NewApp()
	app.Name = "kn"
	app.Usage = "CLI tool for kinfool meta-framework"

	app.Commands = []cli.Command{
		{
			Name:    "initialize",
			Aliases: []string{"i"},
			Usage:   "Create a new kinfool project",
			Action:  func(c *cli.Context) error { initializeAction(); return nil },
		},
		{
			Name:    "tidy",
			Aliases: []string{"t"},
			Usage:   "Tidy up kinfool.go",
			Action:  func(c *cli.Context) error { tidyAction(); return nil },
		},
		{
			Name:    "route",
			Aliases: []string{"r"},
			Usage:   "Create a new route alongside with controller and service",
			Action:  func(c *cli.Context) error { routeAction(c.Args()[0]); return nil },
		},
		{
			Name:    "utility",
			Aliases: []string{"u"},
			Usage:   "Create a new utility",
			Action:  func(c *cli.Context) error { utilityAction(); return nil },
		},
		{
			Name:    "middleware",
			Aliases: []string{"m"},
			Usage:   "Create a new middleware",
			Action:  func(c *cli.Context) error { middlewareAction(c.Args()[0]); return nil },
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
