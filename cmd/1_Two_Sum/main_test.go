package main

import (
	"reflect"
	"sort"
	"testing"
)

var testCases = []struct {
	nums   []int
	target int
	want   []int
}{
	{
		nums:   []int{2, 7, 11, 15},
		target: 9,
		want:   []int{0, 1},
	},
	{
		nums:   []int{3, 2, 4},
		target: 6,
		want:   []int{1, 2},
	},
	{
		nums:   []int{3, 3},
		target: 6,
		want:   []int{0, 1},
	},
}

func TestTwoSum(t *testing.T) {
	t.Run("twoSum1", func(t *testing.T) {
		for _, testCase := range testCases {
			want := testCase.want
			got := twoSum1(testCase.nums, testCase.target)

			if !reflect.DeepEqual(want, got) {
				t.Fatalf("want: %v; got: %v\n", want, got)
			}
		}
	})

	t.Run("twoSum2", func(t *testing.T) {
		for _, testCase := range testCases {
			want := testCase.want
			got := twoSum2(testCase.nums, testCase.target)

			if !reflect.DeepEqual(want, got) {
				t.Fatalf("want: %v; got: %v\n", want, got)
			}
		}
	})

	t.Run("twoSum3", func(t *testing.T) {
		for _, testCase := range testCases {
			want := testCase.want
			got := twoSum3(testCase.nums, testCase.target)
			sort.Ints(got)

			if !reflect.DeepEqual(want, got) {
				t.Fatalf("want: %v; got: %v\n", want, got)
			}
		}
	})
}
