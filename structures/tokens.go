package structures

import "fmt"

type Token struct {
	Name           string
	NumOccurrences int
	LocationFile   string
	LocationLine   []int
}

func (t Token) PrintToken() error {
	if len(t.LocationFile) != len(t.LocationLine) || len(t.LocationFile) != t.NumOccurrences {
		return fmt.Errorf("Token has inconsistent number of occurences")
	}
	fmt.Printf("Token %s with %d occurences at locations:\n", t.Name, t.NumOccurrences)
	for i := range t.LocationFile {
		fmt.Printf("%s at line %d\n", t.LocationFile, t.LocationLine[i])
	}
	return nil
}

func (t *Token) AddOccurence(file string, line int) {
	t.LocationLine = append(t.LocationLine, line)
	t.NumOccurrences++
}

func NewToken(name, file string, numOccurences int, lines []int) Token {
	return Token{
		Name:           name,
		LocationFile:   file,
		NumOccurrences: numOccurences,
		LocationLine:   lines,
	}
}
