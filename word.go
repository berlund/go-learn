package learn

import "strings"

func Word(filename string, word string) (bool, error) {
	text, err := Cat(filename)
	if err != nil {
		return false, err
	}

	contains := strings.Contains(text, word)
	return contains, nil
}
