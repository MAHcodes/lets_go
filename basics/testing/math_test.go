package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			[]int{1, 2, 3}, 6,
		},
		{
			[]int{2, 3, 4}, 9,
		},
		{
			[]int{1, 10, 100}, 111,
		},
	}

	for _, test := range tests {
		sum := Add(test.nums)
		if sum == test.want {
			t.Logf("Add(%v): PASSED, expected %v, got %v", test.nums, test.want, sum)
		} else {
			t.Fatalf(`Add(%v): FAILED, expected %v, got %v`, test.nums, test.want, sum)
		}
	}
}
