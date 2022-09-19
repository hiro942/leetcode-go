package leetcode

import "strconv"

var ans []string

func restoreIpAddresses(s string) []string {
	ans = []string{}
	dfs(s, "", 0, 0)
	return ans
}

func dfs(s, path string, u, cnt int) {
	if u == len(s) {
		if cnt == 4 {
			ans = append(ans, path[:len(path)-1])
		}
		return
	}
	if cnt == 4 {
		// 已经有4个数段但是还没搜完，说明至少该ip地址有五个数段，不合法
		return
	}

	for i, t := u, 0; i < len(s); i++ {
		if i > u && s[u] == '0' {
			break
		}
		t = t*10 + int(s[i]-'0')
		if t <= 255 {
			dfs(s, path+strconv.Itoa(t)+".", i+1, cnt+1)
		} else {
			break
		}
	}
}
