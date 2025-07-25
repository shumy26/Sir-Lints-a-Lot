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
		log.Fatal("No filepath provided")
	}

	path := os.Args[1]
	if !strings.HasSuffix(path, ".py") {
		log.Fatalf("Not a Python file")
	}

	fileBytes, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}
	fileText := string(fileBytes)

	blockList := structures.BlocksFromFile(fileText, path)

	for _, block := range blockList {
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
}
