package main

// 0ms, 2.9MB
func longestValidParentheses(s string) int {
	var top, max, curlen int
	var prelen = make([]int, 0, len(s)) // 预分配内存并初始化长度为0,减少了4ms
	for _, c := range s {
		switch c {
		case '(':
			top++
			prelen, curlen = append(prelen, curlen), 0
		case ')':
			if top > 0 {
				curlen += 2
				top--
				curlen += prelen[top]
				prelen = prelen[:top]
				if curlen > max {
					max = curlen
				}
			} else {
				top, curlen = 0, 0
				prelen = prelen[:0]
			}
		}
	}
	return max
}
