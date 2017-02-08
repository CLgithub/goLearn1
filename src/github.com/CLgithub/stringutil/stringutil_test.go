package stringutil

import "testing"

func Test(t *testing.T) {
	var tests = []struct {
		s, want string
	}{
		{"Abcde","edcbA"},
		{"Hello,世界","界世,olleH"},
		{"",""},
	}
	for _, c := range tests {
		got := Rev(c.s)
		if got != c.want {
			t.Errorf("Rev(%q)==%q, want %q", c.s, got, c.want)
		}
	}
}
