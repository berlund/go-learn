package learn

import "strings"

func MyGrep(filename, word string) ([]int, error) {

	text, err := Cat(filename)
	if err != nil {
		return nil, err
	}

	result := make([]int, 0)

	lines := strings.Split(text, "\n")
	for i, line := range lines {
		if strings.Contains(line, word) {
			result = append(result, i+1)
		}
	}

	return result, nil
}
