package splitspace

// go test -benchmem -bench .

import (
	"regexp"
	"strings"
	"unicode"
)

var (
	gre = regexp.MustCompile("[\\w]+")
)

func splitTrim(in string) []string {
	words := strings.Fields(in) //выделение слов и удаление знаков препинания
	tokens := make([]string, 0, len(words))
	for i := 0; i < len(words); i++ {
		words[i] = strings.ToLower(words[i])
		words[i] = strings.TrimFunc(words[i], func(r rune) bool {
			return !unicode.IsLetter(r)
		})
		if words[i] == "" {
			tokens = append(tokens, words[i])
		}
	}
	return tokens
}

func tokenize(str string) []string {
	re := regexp.MustCompile("[\\w]+")
	tokenPositions := re.FindAllStringIndex(str, -1)
	tokens := make([]string, len(tokenPositions))

	for i, pos := range tokenPositions {
		tokens[i] = strings.ToLower(str[pos[0]:pos[1]])
	}

	return tokens
}

func tokenizeGlobal(str string) []string {
	tokenPositions := gre.FindAllStringIndex(str, -1)
	tokens := make([]string, len(tokenPositions))

	for i, pos := range tokenPositions {
		tokens[i] = strings.ToLower(str[pos[0]:pos[1]])
	}

	return tokens
}
