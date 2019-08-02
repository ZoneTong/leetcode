package main

func isMatch(s string, p string) bool {
	// 暴力递归模式
	// if p == "" {
	// 	return s == ""
	// }

	// first_match := (s != "") && (p[0] == '.' || p[0] == s[0])

	// if len(p) > 1 && p[1] == '*' {
	// 	return isMatch(s, p[2:]) || (first_match && isMatch(s[1:], p))
	// } else {
	// 	return first_match && isMatch(s[1:], p[1:])
	// }

	// 动态规划
	memo := make(map[string]bool)
	return dp(s, p, memo)
}

// 动态规划,中间计算存入备忘录.以空间换时间
func dp(s, p string, memo map[string]bool) bool {
	if p == "" {
		return s == ""
	}

	sp := s + "," + p
	b, ok := memo[sp]
	if ok {
		return b
	}

	first_match := (s != "") && (p[0] == '.' || p[0] == s[0])
	if len(p) > 1 && p[1] == '*' {
		b = dp(s, p[2:], memo) || (first_match && dp(s[1:], p, memo))
	} else {
		b = first_match && dp(s[1:], p[1:], memo)
	}
	memo[sp] = b
	return b
}
