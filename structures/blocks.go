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
	return block
}

func (b *Block) CreateTokens() ([]Token, error) {

	var tokenList []Token

	wordCount, wordLines := b.blockWordMaps()

	for name, occurences := range wordCount { //Iterate over the map
		for wordname, wordLines := range wordLines {
			if name == wordname {
				linesOfWord := wordLines

				token := Token{
					Name:          name,
					NumOccurences: occurences,
					LocationFile:  []string{b.LocationFile}, // Block has only string here, not []string
					LocationLine:  linesOfWord,
				}

				tokenList = append(tokenList, token)
			}
		}

	}

	return tokenList, nil
}

func (b *Block) blockWordMaps() (map[string]int, map[string][]int) { //Helper function for CreateTokens()

	wordCount := make(map[string]int)   // Map to store words and their counts (word -> count)
	wordLines := make(map[string][]int) // Map to store on which lines a word appears (word -> lines)

	lines := strings.Split(b.Code, "\n") //Splits each line of the code

	for lineNum, line := range lines { //Iterates over each line
		words := strings.Fields(line)

		for _, w := range words {
			wordCount[w]++                                 //Increase global word count
			wordLines[w] = append(wordLines[w], lineNum+1) // +1 because we don't have line 0 in most IDEs
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
	level_of_indentation := count / 4
	return level_of_indentation
}
