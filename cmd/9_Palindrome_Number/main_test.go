package main

import (
	"testing"
)

var testCases = []struct {
	x    int
	want bool
}{
	{
		x:    121,
		want: true,
	},
	{
		x:    -121,
		want: false,
	},
	{
		x:    10,
		want: false,
	},
}

func TestPalindromeNumber(t *testing.T) {
	for _, testCase := range testCases {
		want := testCase.want
		got := isPalindrome(testCase.x)

		if want != got {
			t.Fatalf("want: %v; got: %v\n", want, got)
		}
	}
}
