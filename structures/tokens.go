package structures

import "fmt"

type Token struct {
	Name          string
	NumOccurences int
	LocationFile  []string
	LocationLine  []int
}

func (t Token) PrintToken() error {
	if len(t.LocationFile) != len(t.LocationLine) {
		return fmt.Errorf("Token has inconsistent number of occurences")
	}
	fmt.Printf("Token %s with %d occurences at locations:\n", t.Name, t.NumOccurences)
	for i := range t.LocationFile {
		fmt.Printf("%s at line %d\n", t.LocationFile[i], t.LocationLine[i])
	}
	return nil
}

func (t *Token) AddOccurence(file string, line int) {
	t.LocationFile = append(t.LocationFile, file)
	t.LocationLine = append(t.LocationLine, line)
}
