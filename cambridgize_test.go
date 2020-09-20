package main

import (
	"fmt"
	"testing"
)

type testcase struct {
	desc string
	in   string
}

func TestCambridgize(t *testing.T) {
	testcases := []testcase{
		{"regular text", "Hello my name is Greg"},
		{"unchanged", "hey how are you now"},
		{"numbers", "Sure 12345 is a cool number, but I prefer 54321"},
		{"special chars", "Glaçons Über déjà et cætera"},
		{"irregular separators", "Hello aujourd'hui c'est tourne-disque"},
		{"punctuation", "heyyy???? Help!!!! I'm Kev,,,, Kev Adams...."},
	}
	// rgx := wordRegexp()

	// TODO: actual ckeck
	for _, tc := range testcases {
		fmt.Println(Cambridgize(tc.in))
	}

}

func isValidWordOutput(in, out string) bool {
	n, k := len(in), len(out)

	if n != k {
		return false
	}

	if n <= 3 {
		return in == out
	}

	return in[0] == out[0] && in[n-1] == out[n-1]
}
