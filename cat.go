package learn

import (
	"fmt"
	"io/ioutil"
)

func Cat(filename string) (string, error) {

	body, err := ioutil.ReadFile(filename)

	if err != nil {
		return "", fmt.Errorf("File %q not found", filename)
	}

	s := string(body)
	fmt.Println(s)
	return s, nil
}
