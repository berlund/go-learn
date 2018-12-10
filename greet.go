package learn

import (
	"fmt"
)

func Greet(language, who string) string {
	hi := "Hi"
	if language == "sv" {
		hi = "Hej"
	}

	format := "%s, %s!"
	return fmt.Sprintf(format, hi, who)
}
