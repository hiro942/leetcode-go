package leetcode

import (
	"sort"
)

func groupAnagrams(strs []string) (ans [][]string) {
	hash := map[string][]string{}
	for _, s := range strs {
		t := sortString(s)
		hash[t] = append(hash[t], s)
	}

	for k := range hash {
		ans = append(ans, hash[k])
	}

	return ans
}

func sortString(s string) string {
	t := s
	sort.Slice(t, func(i, j int) bool {
		return t[i] < t[j]
	})
	return t
}
