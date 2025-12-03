package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func checkKinfool() bool {
	if _, err := os.Stat("./kinfool.go"); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Println("kinfool.go does not exist in the current directory")
		os.Exit(1)
	}
	return false
}

func updateKinfool(content string) bool {
	if file, err := os.Create("./kinfool.go"); err == nil {
		file.WriteString(content)
		file.Close()
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Println("Error writing to file:", err)
		return false
	}
	return false
}

func getMainModuleName() string {
	if data, err := os.ReadFile("./go.mod"); err == nil {
		return string(data)
	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Println("go.mod does not exist in the current directory")
		os.Exit(1)
	}
	return ""
}

func checkArg(c string, m string) bool {
	if len(c) < 1 {
		fmt.Printf("Missing %s", m)
		os.Exit(1)
	}
	return true
}

func fileExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Println(err)
		os.Exit(1)
	}
	return false
}

func getDirectoryContents(path string) []string {
	if entries, err := os.ReadDir(path); err == nil {
		var fileNames []string // Initialize an empty string slice
		for _, entry := range entries {
			fileNames = append(fileNames, entry.Name())
		}
		return fileNames
	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Println(err)
		return []string{}
	}
	return []string{}
}

func createNewFile(filepath string, content string) bool {
	if file, err := os.Create(filepath); err == nil {
		file.WriteString(content)
		file.Close()
		return true
	} else if errors.Is(err, os.ErrNotExist) {
		fmt.Println("Error writing to file:", err)
		return false
	}
	return false
}

func createNewFolder(folderpath string) bool {
	if err := os.MkdirAll(folderpath, os.ModeAppend); err == nil {
		return true
	} else {
		fmt.Println("Error creating folder:", err)
		return false
	}
}

func runCmd(main string, params ...string) bool {
	cmd := exec.Command(main, params...)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing command:", err)
		return false
	}
	fmt.Println(string(output))
	return true
}
