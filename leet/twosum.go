package leet

func twoSum(nums []int, target int) []int {
	mapping := map[int]int{}

	for idx, value := range nums {
		partner := target-value

		if val, exists := mapping[partner]; exists {
			return []int{val, idx}
		}
		mapping[value] = idx 
	}
	return nil
}
