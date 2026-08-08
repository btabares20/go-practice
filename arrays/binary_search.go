package arrays

import "math"

func binary_search(haystack []int, needle int) bool{
	var lo int = 0;
	var hi int = len(haystack);
	
	for {
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
		} else {
			return false
		}
	}
}
