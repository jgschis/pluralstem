package pluralstem

import (
	"fmt"

	"github.com/blevesearch/pluralstem/english"
)

func Stem(word string, language string) (string, error) {

	var f func(string) string

	switch language {
	case "english":
		f = english.Stem
	default:
		err := fmt.Errorf("unknown language: %v", language)
		return "", err
	}

	stemmed := f(word)
	return stemmed, nil
}
