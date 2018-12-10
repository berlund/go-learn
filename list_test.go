package learn

import (
	"testing"
)

type weekday uint

// or use iota here
const (
	Monday   weekday = 0
	Tuesday  weekday = 1
	Saturday weekday = 5
	Sunday   weekday = 6
)

func TestSlice(t *testing.T) {
	from := Saturday
	got := Day(from, 6) // 0..6 []string weekdays
	a, b := "Saturday", "Sunday"
	if a != got[0] || b != got[1] {
		t.Fail()
	}
}

func Day(from, to weekday) []string {
	week := make([]string, 7)
	week[0] = "Monday"
	week[5] = "Saturday"
	week[6] = "Sunday"

	return week[from : to+1]
}
