package main

import (
	"reflect"
	"testing"
)

func TestGeneratePatterns(t *testing.T) {
	cases := []struct {
		in   string
		want []string
	}{
		{in: "", want: nil},
		{in: "1", want: nil},
		{in: "11", want: []string{"1"}},
		{in: "22", want: []string{"2"}},
		{in: "99", want: []string{"9"}},
		{in: "121", want: []string{"1"}},
		{in: "1010", want: []string{"1", "10"}},
		{in: "222222", want: []string{"2", "22", "222"}},
		{in: "1188511885", want: []string{"1", "11", "1188", "11885"}},
	}

	for _, c := range cases {
		got := generatePatterns(c.in)
		if !reflect.DeepEqual(got, c.want) {
			t.Fatalf("generatePatterns(%q) = %v, want %v", c.in, got, c.want)
		}
	}
}
