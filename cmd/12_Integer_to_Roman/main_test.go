package main

import "testing"

var testCases = []struct {
	num  int
	want string
}{
	{
		num:  3749,
		want: "MMMDCCXLIX",
	},
	{
		num:  58,
		want: "LVIII",
	},
	{
		num:  1994,
		want: "MCMXCIV",
	},
}

func TestIntToRoman(t *testing.T) {
	for _, testCase := range testCases {
		want := testCase.want
		got := intToRoman(testCase.num)

		if want != got {
			t.Fatalf("want: %v; got: %v\n", want, got)
		}
	}
}
