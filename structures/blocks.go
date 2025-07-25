package structures

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

func (b *Block) CreateTokens() []Token, error { //DONT FORGET TO IMPORT "strings"

	wordList, numOccurences := blockWordList(b *Block)

	return nil
}

func blockWordList(b *Block) []string, []int { //Helper function for CreateTokens()
	 wordCount := make(map[string]int) // Map to store words and their counts

    lines := strings.Split(b.Code, "\n")

    for _, line := range lines {
        words := strings.Fields(line) //Separate words on each line
        for _, w := range words {
            wordCount[w]++
        }
    }

    var wordList []string
    var numOccurences []int

    for word, count := range wordCount {
        wordList = append(wordList, word)
        numOccurences = append(numOccurences, count)
    }
	
    return wordList, numOccurences
}

func countLeadingWhitespace(line string) int { //Helper function to count the leading whitespace on each line of the code
    count := 0
    for _, ch := range line {
        if ch == ' ' || ch == '\t' {
            count++
        } else {
            break
        }
    }
    return count
}

/* type Token struct {
	Name          string
	NumOccurences int
	LocationFile  []string
	LocationLine  []int
} */