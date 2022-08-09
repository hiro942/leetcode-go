package leetcode

func isValid(s string) bool {
	matches := map[byte]byte{'[': ']', '{': '}', '(': ')'}
	var stack []byte
	for i := range s {
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		} else {
			if len(stack) > 0 && matches[stack[len(stack)-1]] == s[i] {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}
