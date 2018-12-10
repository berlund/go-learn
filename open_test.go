package learn

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestOpenFile(t *testing.T) {
	_, err := os.Open("nosuchfile.txt")
	if err == nil {
		t.Fail()
	}

	file, err := os.Open("open_test.go")
	if err != nil {
		t.Fail()
	}

	body, err := ioutil.ReadAll(file)
	if err != nil || len(body) == 0 {
		t.Fail()
	}

	file.Close()
}

func TestCat(t *testing.T) {
	_, err := Cat("nothing")
	if err == nil {
		t.Fail()
	}

	text, err := Cat("test.txt")
	if err != nil {
		t.Fail()
	}

	if text != "hello\nworld\n" {
		t.Fail()
	}
}

func TestWord(t *testing.T) {
	found, err := Word("test.txt", "world")
	if err != nil {
		t.Fail()
	}

	if !found {
		t.Errorf("Word not found")
	}

	found, _ = Word("test.txt", "foo")
	if found {
		t.Errorf("Found non existing word")
	}

}

func TestGrep(t *testing.T) {
	lines, _ := MyGrep("test.txt", "world")
	if len(lines) != 1 || lines[0] != 2 {
		t.Errorf("Wrong line")
	}

	lines, _ = MyGrep("test2.txt", "bar")
	t.Logf("%v", lines)
	if len(lines) != 3 || lines[0] != 2 || lines[1] != 3 || lines[2] != 5 {
		t.Errorf("Wrong line")
	}
}
