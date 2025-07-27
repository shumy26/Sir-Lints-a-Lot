package structures

import {
	"fmt"
	"errors"
	}

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

func analyzeCode(t *GlobalTokenMap) ([]string, error) {
	var problems []string

	if len(t.TokensWithProblems) > 0 {
		for _, value := range t.TokensWithProblems {
		
		errorString := fmt.Sprintf("Variable %s was declared on line %d but wasn't used!", value.Name, value.LocationLine[0])
		problems = append(problems, errorString)
	}
	}else{
		return nil, errors.New("No problematic tokens found!")
	}

	return problems, nil
}
