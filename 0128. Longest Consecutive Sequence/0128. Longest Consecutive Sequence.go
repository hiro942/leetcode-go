package leetcode

func longestConsecutive(nums []int) int {
	res := 0
	set := make(map[int]struct{})
	for _, v := range nums {
		set[v] = struct{}{}
	}
	for _, v := range nums {
		_, ok1 := set[v]
		_, ok2 := set[v-1]
		if ok1 && !ok2 {
			u := v
			for {
				if _, ok := set[u]; ok {
					delete(set, u)
					u++
				} else {
					break
				}
			}
			if u-v > res {
				res = u - v
			}
		}
	}
	return res
}
