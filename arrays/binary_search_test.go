package arrays

import "testing"

func TestBinarySearch(t *testing.T) {
	t.Run("binary_searc", func(t *testing.T) {
		haystack := [9]int{1,2,3,4,5,7,8,9,10}
		if binary_search(haystack[:], 9)	== false {
			t.Errorf("should be true")
		}
	});
}
