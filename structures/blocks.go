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

	wordList, numOccurences := blockWordMaps(b *Block)
	
	for i := 0; i < len(wordList); i++ {

		token := Token{
			Name:          wordList[i]
			NumOccurences: numOccurences[i]
			LocationFile  []string
			LocationLine  []int

		}

	}

	return nil
}

func blockWordMaps(b *Block) map[string]int,  { //Helper function for CreateTokens()

	wordCount := make(map[string]int) 	 	// Map to store words and their counts (word -> count)
	wordLineCount := make(map[string][]int)  // Map to store on which lines a word appears (word -> lines)
    wordLineOccur := make(map[string][]int)   // Map to store num of occurences per line of a word (word -> occurences)

    lines := strings.Split(b.Code, "\n") //Splits each line of the code

 	for lineNum, line := range lines { //Iterates over each line
        words := strings.Fields(line)
        lineWordCount := make(map[string]int)

        for _, w := range words {
            wordCount[w]++ //Increase global word count
            lineWordCount[w]++ //Word count JUST for that particular line
        }

        for w, count := range lineWordCount {
            wordLineCount[w] = append(wordLineCount[w], lineNum+1) // +1 because we don't have line 0 in most IDEs
            wordLineOccur[w] = append(wordLineOccur[w], count)
        }
    }

    return wordCount, wordLineCount, wordLineOccur
}

func countLeadingWhitespace(line string) int { //Helper function to count the leading whitespace on each line of the code
    count := 0
    for _, ch := range line {
        if ch == ' ' || ch == '\t' { //Separates by leading whitespace or "tabs"
            count++
        } else {
            break
        }
    }
    return count
}
