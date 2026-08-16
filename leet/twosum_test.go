package leet

import (
	"reflect"
	"testing"
)


func TestTwoSum(t *testing.T) {
	test_cases := []struct{
		nums []int
		want []int
		target int
	}{
		{nums:[]int{2,7,11,15}, want:[]int{0,1}, target:9},
		{nums:[]int{3,2,4}, want:[]int{1,2}, target:6},
		{nums:[]int{3,3},want:[]int{0,1}, target:6},
	}
	for _, tt := range test_cases {
		got := twoSum(tt.nums, tt.target)
		assertMessage(t, got, tt.want)
	}
}

func assertMessage(t testing.TB, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
