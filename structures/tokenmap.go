package structures

import (
	"fmt"
	"math/rand"
)

const bright = ";1"
const underline = ";4"
const red = "\033[31"
const yellow = "\033[33"
const blue = "\033[34"
const green = "\033[32"
const reset = "\033[0m"

type GlobalTokenMap struct {
	TokenMap           map[string]Token // key is token name
	TokensWithProblems map[string]Token
}

func (t *GlobalTokenMap) AddToken(token Token) error {
	if _, ok := t.TokenMap[token.Name]; ok {
		return fmt.Errorf("token with this name already exists")
	} else if token.NumOccurrences <= 1 {
		t.TokenMap[token.Name] = token
		t.TokensWithProblems[token.Name] = token
	} else {
		t.TokenMap[token.Name] = token
	}
	return nil
}

func AnalyzeCode(t *GlobalTokenMap) ([]string, error) {
	var problems []string
	if len(t.TokensWithProblems) > 0 {
		fmt.Printf("%v%v%vmSir Monty Lints-a-Lot has analyzed your code%v\n", blue, bright, underline, reset)
		for _, value := range t.TokensWithProblems {
			errorString := fmt.Sprintf(flavourGenerator(), yellow, bright, value.Name, reset, yellow, bright, value.LocationLine[0], reset)
			problems = append(problems, errorString)
		}
	} else {
		return nil, fmt.Errorf("%v%vmNo problematic tokens found!%v", green, bright, reset)
	}

	return problems, nil
}

func flavourGenerator() string {
	flavour := []string{
		"What is the air-speed velocity of a swallow ladened by declaring %v%vm%s%v on line %v%vm%d%v and not using it?\n",
		"What makes you think she's a witch? Well, she declared %v%vm%s%v on line %v%vm%d%v and didn't use it.\n",
		"Your mother was a hamster, and your father declared %v%vm%s%v on line %v%vm%d%v and didn't use it.\n",
		"Armaments, chapter two, verses nine through twenty-one. And the Lord spake, saying, First shalt thou declare %v%vm%s%v on line %v%vm%d%v and then shalt thou use it.\n",
		"Look you stupid bastard you've declared %v%vm%s%v on line %v%vm%d%v and haven't used it. Yes I have, Look! It's just a flesh wound.\n",
		"Come and see the violence inherent in the system. Help! Help! I'm being forced to declare %v%vm%s%v on line %v%vm%d%v without using it!\n",
		"Ni! you must go to the enchanted forest and declare the largest variable named %v%vm%s%v with a herring on line %v%vm%d%v, and then use it. Ni, Ni, Ni!\n",
	}
	return flavour[rand.Intn(len(flavour))]
}
