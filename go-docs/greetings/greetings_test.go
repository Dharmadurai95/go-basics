package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	name := "Dharma"
	want := regexp.MustCompile(`\b` + name + `\b`)
	message, err := Hello(name)
	if !want.MatchString(message) || err != nil {
		t.Fatalf(`Hello ("Dharma") ,%s ,%v want match for %#q,nil`, message, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {

	name := ""
	msg, err := Hello(name)
	if err != nil {
		t.Fatalf(`Hello("") = %s, %v, want "", error`, msg, err)
	}
}
