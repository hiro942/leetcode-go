package leetcode

func singleNumber(nums []int) int {
	res := 0
	for bit := 0; bit < 32; bit++ {
		cnt := 0
		for _, v := range nums {
			cnt += (v >> bit) & 1
		}
		if cnt%3 == 1 {
			res += 1 << bit
		}
	}
	return int(int32(res)) // 可能会溢出int32
}
