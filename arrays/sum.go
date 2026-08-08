package arrays



func Sum(nums []int) int {
	var total int = 0
	for _, v :=range nums {
		total += v
	}
	return total
}
 func SumAll(nums ...[]int) []int{
	 var total []int;
	 for _, v :=range nums {
		 total  = append(total, Sum(v))
	 }
	 return total
 }
 func SumAllTails(nums ...[]int) []int{
	 var total []int;
	 for _, v :=range nums {
		 if len(v) > 0 {
			 total  = append(total, Sum(v[1:]))
		 } else {
		 	total = append(total, Sum(v))
		}
	 }
	 return total
 }

