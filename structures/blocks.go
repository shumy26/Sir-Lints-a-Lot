package structures

import (
	"strings"
)

type Block struct {
	Code              string
	LocationFile      string
	LocationLineStart int
	LocationLineEnd   int
	TokenList         []Token
}

func CreateBlock(code, file string, lineStart, lineEnd int) Block {
	block := Block{
		Code:              code,
		LocationFile:      file,
		LocationLineStart: lineStart,
		LocationLineEnd:   lineEnd,
	}
	block.CreateTokens()
	return block
}

func (b *Block) CreateTokens() {

	var tokenList []Token

	wordCount, wordLines := b.blockWordMaps()

	for name, occurrences := range wordCount { //Iterate over the map
		for wordname, wordLines := range wordLines {
			if name == wordname {
				token := NewToken(name, b.LocationFile, occurrences, wordLines) //function from tokens.go
				tokenList = append(tokenList, token)                            //append to a list of all Tokens on this block
			}
		}
	}
	b.TokenList = append(b.TokenList, tokenList...)
}

func (b *Block) blockWordMaps() (map[string]int, map[string][]int) { //Helper function for CreateTokens()

	pythonKeywords := [35]string{"False", "await", "else", "import", "pass", "None", "break", "except", "in", "raise", "True", "class", "finally", "is", "return", "and", "continue", "for", "lambda", "try", "as", "def", "from", "nonlocal", "while", "assert", "del", "global", "not", "with", "async", "elif", "if", "or", "yield"}

	wordCount := make(map[string]int)   // Map to store words and their counts (word -> count)
	wordLines := make(map[string][]int) // Map to store on which lines a word appears (word -> lines)

	lines := strings.Split(b.Code, "\n") //Splits each line of the code

	for lineNum, line := range lines { //Iterates over each line
		trimmed := strings.TrimSpace(line)    //Removes whitespace before "#" - we'll consider it for the indentation in another function
		if !strings.HasPrefix(trimmed, "#") { //Completely ignores line if it starts with # (it is a comment!)

			words := strings.Fields(line)

			for _, w := range words {
				isKeyword := false
				for _, p := range pythonKeywords { //Quick loop to check if the word is a Python keyword, in which case it should be ignored
					if w == p {
						isKeyword = true
						break
					}
				}
				if !isKeyword {
					wordCount[w]++                                 //Increase word count for this block
					wordLines[w] = append(wordLines[w], lineNum+1) // +1 because we don't have line 0 in most IDEs
				}
			}
		}
	}
	return wordCount, wordLines
}

func countLeadingWhitespace(line string) int { //Helper function to count the leading whitespace on each line of the code
	count := 0
loop:
	for _, ch := range line {
		switch ch {
		case ' ': //Separates by leading whitespace or "tabs"
			count++
		case '\t':
			count += 4 //Assuming a tab is equivalent to 4 spaces.
		default:
			break loop
		}
	}
	return count
}

func BlocksFromFile(fileText, fileName string) []Block {
	blockList := make([]Block, 0)

	lines := strings.Split(fileText, "\n")

	indentationList := make([]int, 0)
	for _, line := range lines {
		indentationLevel := countLeadingWhitespace(line)
		indentationList = append(indentationList, indentationLevel)
	}

	for idx := 0; idx < len(indentationList); idx++ {
		endIdx := idx
		if idx == 0 {
			//global scope, block with the whole code
			endIdx = len(indentationList)
			code := strings.Join(lines[idx:endIdx], "\n")
			blockList = append(blockList, CreateBlock(
				code,
				fileName,
				idx+1,
				endIdx))
		} else {
			if indentationList[idx] > indentationList[idx-1] {
				for j := idx + 1; j < len(indentationList) && indentationList[j] >= indentationList[idx]; j++ {
					endIdx = j
				}
				code := strings.Join(lines[idx:endIdx+1], "\n")
				blockList = append(blockList, CreateBlock(
					code,
					fileName,
					idx+1,
					endIdx+1,
				))
			}
		}
	}
	return blockList
}
