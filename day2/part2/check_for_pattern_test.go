package main

import "testing"

func TestCheckForPattern(t *testing.T) {
	cases := []struct {
		s       string
		start   int
		pattern string
		want    int
	}{
		{s: "11", start: 0, pattern: "1", want: 0},
		{s: "11", start: 1, pattern: "1", want: 1},
		{s: "1212012", start: 0, pattern: "12", want: 0},
		{s: "1212012", start: 1, pattern: "12", want: 2},
		{s: "1212012", start: 3, pattern: "12", want: 5},
		{s: "1212012", start: 0, pattern: "99", want: -1},
		{s: "1299012", start: 0, pattern: "99", want: 2},
		{s: "1299992", start: 4, pattern: "99", want: 4},
		{s: "123", start: 0, pattern: "123", want: 0},
		{s: "123", start: 1, pattern: "123", want: -1},
		{s: "abcde", start: 3, pattern: "de", want: 3},
		{s: "", start: 0, pattern: "a", want: -1},
	}

	for _, c := range cases {
		got := checkForPattern(c.s, c.start, c.pattern)
		if got != c.want {
			t.Fatalf("checkForPattern(%q, %d, %q) = %d, want %d", c.s, c.start, c.pattern, got, c.want)
		}
	}
}
