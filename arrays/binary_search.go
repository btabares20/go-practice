package arrays

import "math"

func search(haystack []int, lo, hi, needle int) bool{
	var mid int = int(math.Floor(
		float64(lo+(hi-lo)/2),
	));
	if lo < hi {
		if haystack[mid] == needle {
			return true
		}

		if haystack[mid] < needle {
			lo = mid + 1
		} else {
			hi = mid
		}
		return search(haystack, lo, hi, needle)
	} else {
		return false
	}
}
func binary_search(haystack []int, needle int) bool{
	var lo int = 0;
	var hi int = len(haystack);

	return search(haystack, lo, hi, needle)
}
