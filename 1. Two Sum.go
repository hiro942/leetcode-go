package leetcode

func twoSum(nums []int, target int) []int {
	var m = map[int]int{} // num -> its index

	for i, val := range nums {
		if _, ok := m[target-val]; ok {
			return []int{m[target-val], i}
		}
		m[val] = i
	}

	return []int{}
}
