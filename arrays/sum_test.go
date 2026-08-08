package arrays

import (
	"reflect"
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	assertMessage := func(t testing.TB, got, want int) {
		t.Helper()
		if got!=want{
			t.Errorf("got %d want %d", got, want)
		}
	}
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
func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("make the sums of tails of", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})

}
