package main

import (
	"testing"
)

var testCases = []struct {
	strs []string
	want string
}{
	{
		strs: []string{"flower", "flow", "flight"},
		want: "fl",
	},
	{
		strs: []string{"dog", "racecar", "car"},
		want: "",
	},
}

func TestLongestCommonPrefix(t *testing.T) {
	for _, testCase := range testCases {
		want := testCase.want
		got := longestCommonPrefix(testCase.strs)

		if want != got {
			t.Fatalf("want: %v; got: %v\n", want, got)
		}
	}
}
