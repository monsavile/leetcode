package main

import (
	"testing"
)

var testCases = []struct {
	s    string
	want int
}{
	{
		s:    "III",
		want: 3,
	},
	{
		s:    "LVIII",
		want: 58,
	},
	{
		s:    "MCMXCIV",
		want: 1994,
	},
}

func TestRomanToInteger(t *testing.T) {
	t.Run("romanToInt", func(t *testing.T) {
		for _, testCase := range testCases {
			want := testCase.want
			got := romanToInt(testCase.s)

			if want != got {
				t.Fatalf("want: %v; got: %v\n", want, got)
			}
		}
	})
}
