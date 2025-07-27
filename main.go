package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/shumy26/Sir-Lints-a-Lot/structures"
)

var globalTokenMap = &structures.GlobalTokenMap{
	TokenMap:           make(map[string]structures.Token),
	TokensWithProblems: make(map[string]structures.Token),
}

func main() {
	if len(os.Args) < 2 {
		var files []string
		err := grabFiles(".", &files)
		if err != nil {
			log.Fatalf("Error grabbing files: %v", err)
		}
		fmt.Println("Files found in the directory:")
		for _, file := range files {
			fmt.Println(file)
		}
		log.Print("Please enter a Python file path from this list.")
	}

	var input string
	fmt.Scanln(&input)

	path := input
	if !strings.HasSuffix(path, ".py") {
		log.Fatalf("Not a Python file")
	}

	fileBytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	fileText := string(fileBytes)

	blockList := structures.BlocksFromFile(fileText, path)

	fmt.Println("Please choose the scope you want to inspect:")
	for idx, block := range blockList {
		if idx == 0 {
			fmt.Printf("Global Scope (0) from lines: <-- %d : %d -->\n", block.LocationLineStart, block.LocationLineEnd)
		} else {
			fmt.Printf("Scope %d from lines : <-- %d : %d -->\n", idx, block.LocationLineStart, block.LocationLineEnd)
		}

	}

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter your choice (scope number, 0 for global): ")
		input, _ := reader.ReadString('\n')
		inputStr := strings.TrimSpace(input)
		inputInt, err := strconv.Atoi(inputStr)
		if err != nil {
			log.Fatal("Invalid input, please choose a number")
		}

		if inputInt >= len(blockList) {
			fmt.Printf("Invalid choice %d, please enter a valid number\n", inputInt)
		} else {
			block := blockList[inputInt]
			for i := 0; i < len(block.TokenList); i++ {
				globalTokenMap.AddToken(block.TokenList[i])
			}
			//fmt.Println(globalTokenMap)
			break
		}
	}
}
