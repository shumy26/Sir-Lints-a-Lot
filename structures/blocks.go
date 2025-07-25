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

func (b *Block) CreateTokens() error {
	return nil
}
