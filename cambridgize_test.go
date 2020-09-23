package main

import (
	"regexp"
	"testing"
)

func TestCambridgize(t *testing.T) {
	testcases := []struct {
		desc string
		in   string
		rgx  string
	}{
		{
			desc: "regular text",
			in:   "Hello my name is Greg",
			rgx:  `^H[el]{3}o my n[am]{2}e is G[re]{2}g$`,
		}, {
			"no change",
			"hey how are you now",
			`^hey how are you now$`,
		}, {
			"numbers",
			"Sure 12345 is a cool number, but I prefer 54321",
			`^S[ur]{2}e 12345 is a c(o){2}l n[umbe]{4}r, but I p[ref]{4}r 54321$`,
		}, {
			"special chars",
			"Glaçons Über déjà et cætera",
			`^G[laçon]{5}s Ü[be]{2}r d[éj]{2}à et c[æter]{4}a$`,
		}, {
			"irregular separators",
			"Hello aujourd'hui c'est tourne-disque",
			`^H[el]{3}o a[ujor]{5}d'hui c'est t[ourn]{4}e-d[isqu]{4}e$`,
		}, {
			"punctuation",
			"heyyy???? Help!!!! I'm Kev,,,, Kev Adams....",
			`^h[ey]{3}y\?\?\?\? H[el]{2}p!!!! I'm Kev,,,, Kev A[dam]{3}s\.\.\.\.$`,
		},
	}

	for _, tc := range testcases {
		rgx := regexp.MustCompile(tc.rgx)
		got := Cambridgize(tc.in)

		if !rgx.MatchString(tc.in) {
			t.Errorf("%s: invalid output expected %s, got %s\n", tc.desc, tc.rgx, got)
		}

		if got == tc.in && rgx.MatchString("/w{4,}") {
			t.Errorf("%s: expected changes, found no changes", tc.desc)
		}
	}
}

func TestCambridgizeWord(t *testing.T) {
	testcases := []struct {
		desc string
		in   string
		exp  []string
	}{
		{"no change 1 letter", "a", []string{"a"}},
		{"no change 2 letters", "in", []string{"in"}},
		{"no change 3 letters", "out", []string{"out"}},
		{"regular word", "Gophr", []string{"Gophr", "Gohpr", "Ghopr", "Ghpor", "Gphor", "Gpohr"}},
		{"special char", "àπŸß", []string{"àπŸß", "àŸπß"}},
	}

	for _, tc := range testcases {
		got := cambridgizeWord(tc.in)

		for i := 0; i < 10; i++ {
			if !contains(tc.exp, got) {
				t.Errorf("invalid word output: expected %v, got %s\n", tc.exp, got)
			}
		}
	}
}

func contains(strs []string, s string) bool {
	for _, str := range strs {
		if str == s {
			return true
		}
	}
	return false
}
