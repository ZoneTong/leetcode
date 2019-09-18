package main

// 解法1: 48ms,5.2M. 直接滑动窗口,一段段比较单词
// func findSubstring(s string, words []string) []int {
// 	if len(words) == 0 {
// 		return nil
// 	}

// 	comblen := len(words[0]) * len(words)
// 	// if len(s) < comblen {
// 	// 	return nil
// 	// }

// 	var result []int
// 	start_index, end_index := 0, comblen
// 	for end_index <= len(s) {
// 		if verifySubstrings(s[start_index:end_index], words) {
// 			result = append(result, start_index)
// 		}
// 		start_index++
// 		end_index++
// 	}

// 	return result
// }

// func verifySubstrings(s string, words []string) bool {
// 	useds := make([]bool, len(words))
// 	for s != "" {
// 		found := false
// 		for i, word := range words {
// 			if useds[i] {
// 				continue
// 			}

// 			if s[:len(word)] == word {
// 				useds[i] = true
// 				s = s[len(word):]
// 				found = true
// 			}
// 		}
// 		if !found {
// 			return false
// 		}
// 	}

// 	return true
// }

// 解法2: 4ms,3.2M. 使单词个数一致
func findSubstring(s string, words []string) []int {
	var result []int
	var wslen = len(words)
	if wslen == 0 {
		return result
	}

	comblen := len(words[0]) * wslen
	if len(s) < comblen {
		return result
	}

	counter := make(map[string]int, wslen)
	for _, w := range words {
		counter[w]++
	}

	wlen := len(words[0])
	count := 1
	remain := make(map[string]int, len(counter))
	reset := func() {
		if count == 0 {
			return
		}

		for k, v := range counter {
			remain[k] = v
		}
		count = 0
	}

	left, right := 0, 0
	leftStep := func() {
		remain[s[left:left+wlen]]++
		left += wlen
		count--
	}

	for i := 0; i < wlen; i++ {
		left, right = i, i
		reset()

		for left+comblen <= len(s) {
			remaincnt, ok := remain[s[right:right+wlen]]
			if !ok {
				left, right = right+wlen, right+wlen
				reset()
			} else if remaincnt == 0 {
				leftStep()
			} else {
				remain[s[right:right+wlen]]--
				right += wlen
				count++
				if count == wslen {
					result = append(result, left)
					leftStep()
				}
			}
		}
	}

	return result
}
