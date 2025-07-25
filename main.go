package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

//var globalTokenMap = &structures.GlobalTokenMap{
//	TokenMap:           make(map[string]structures.Token),
//	TokensWithProblems: make(map[string]structures.Token),
//}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("No file found")
	}
	path := os.Args[1]
	if !strings.HasSuffix(path, ".py") {
		fmt.Println("Not a Python file")
		os.Exit(1)
	}
	fileBytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	fileText := string(fileBytes)
	fmt.Println(fileText)
}
