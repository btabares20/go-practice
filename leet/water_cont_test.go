package leet

import "testing"

func TestWater(t *testing.T) {
	test_cases := []struct{
		nums []int
		want int
	}{
		{nums:[]int{1,8,6,2,5,4,8,3,7}, want: 49},
		{nums:[]int{1,1}, want: 1},
	}
	for _, tt := range test_cases {
		got := maxArea(tt.nums)
		assertEqual(t, got, tt.want)
	}
}

func assertEqual(t testing.TB, got, want int) {
	t.Helper()
	if got!= want {
		t.Errorf("got %v want %v", got, want)
	}
}
