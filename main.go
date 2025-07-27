package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/shumy26/Sir-Lints-a-Lot/structures"
)

//var globalTokenMap = &structures.GlobalTokenMap{
//	TokenMap:           make(map[string]structures.Token),
//	TokensWithProblems: make(map[string]structures.Token),
//}

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

	_ = blockList

	// Testing block, uncomment to see the output:

	/*for _, block := range blockList {

		fmt.Println(block, " ")
		for i := 0; i < len(block.TokenList); i++ {
			fmt.Println(" ")
			fmt.Printf("Name:\t%v\n", block.TokenList[i].Name)
			fmt.Printf("Occurrences:\t%v\n", block.TokenList[i].NumOccurrences)
			fmt.Printf("Line Number:\t%v\n", block.TokenList[i].LocationLine)
			fmt.Printf("File:\t%v\n", block.TokenList[i].LocationFile)
			fmt.Println(" ")
		}
	}
	fmt.Println("Blocks found in the file:")
	fmt.Println(len(blockList))
	for _, block := range blockList {
		fmt.Printf("Block code:\n%s\n", block.Code)
		fmt.Printf("Start Line: %d, End Line: %d\n", block.LocationLineStart, block.LocationLineEnd)
		fmt.Println()
	}*/
}
