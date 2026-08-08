package arrays

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("array", func(t *testing.T) {
		var nums = [5]int{1,2,5,6,7}
		var got int = Sum(nums[:])
		var want int = 21
		assertMessage(t, got, want)
	})
	t.Run("slice", func(t *testing.T) {
		var nums = []int{1,2,5,6,7}
		var got int = Sum(nums)
		var want int = 21
		assertMessage(t, got, want)
	})
}
func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
func assertMessage(t testing.TB, got, want int) {
	t.Helper()
	if got!=want{
		t.Errorf("got %d want %d", got, want)
	}
}
