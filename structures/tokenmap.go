package structures

import "fmt"

type GlobalTokenMap struct {
	TokenMap           map[string]Token // key is token name
	TokensWithProblems map[string]Token
}

func (t *GlobalTokenMap) AddToken(token Token) error {
	if _, ok := t.TokenMap[token.Name]; ok {
		return fmt.Errorf("token with this name already exists")
	} else if token.NumOccurences <= 1 {
		t.TokenMap[token.Name] = token
		t.TokensWithProblems[token.Name] = token
	} else {
		t.TokenMap[token.Name] = token
	}
	return nil
}
