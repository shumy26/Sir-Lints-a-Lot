package main

import (
	"fmt"
	"os"
	"strings"
)

//var globalTokenMap = &structures.GlobalTokenMap{
//	TokenMap:           make(map[string]structures.Token),
//	TokensWithProblems: make(map[string]structures.Token),
//}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("No file found")
		os.Exit(1)
	}
	path := os.Args[1]
	if !strings.HasSuffix(path, ".py") {
		fmt.Println("Not a Python file")
		os.Exit(1)
	}
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(1)
	}
	fileText := string(fileBytes)
	fmt.Println(fileText)
}
