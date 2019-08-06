package main

import (
	"bytes"
	"strings"
)

func nearestPalindromic(n string) string {
	half := len(n) / 2
	if half < 1 {
		return string(n[half] - 1)
	}

	var p string
	if !isPalindromic(n) {
		// 后半部分对齐即可快速找出一个回文数
		//		half := len(n) / 2
		if len(n)%2 == 1 {
			p = makePalindromic(n[:half], string(n[half]))
		} else {
			p = makePalindromic(n[:half], "")
		}
	} else {
		p = n
	}
	small, big := PalindromOfPalindrom(p)
	//	println(small, p, big, n)
	var arr, arr1 []string
	if p != n {
		arr = []string{small, p, big}
		arr1 = []string{AbSub(small, n), AbSub(p, n), AbSub(big, n)}
	} else {
		arr = []string{small, big}
		arr1 = []string{AbSub(small, n), AbSub(big, n)}
	}
	return arr[smallest(arr1)]

}

func isPalindromic(n string) bool {
	length := len(n)
	for i := 0; i < length-1-i; i++ {
		if n[i] != n[length-i-1] {
			return false
		}
	}
	return true
}

func makePalindromic(prepart string, middle string) string {
	//	println(prepart, middle)
	buf := bytes.NewBufferString(prepart)
	buf.WriteString(middle)
	for i := range prepart {
		buf.WriteByte(prepart[len(prepart)-1-i])
	}
	return buf.String()
}

func Inc(n string) string {
	length := len(n)
	buf := []byte(n)
	var i int
	for i = length - 1; i >= 0; i-- {
		if buf[i] == '9' {
			buf[i] = '0'
		} else {
			buf[i] = buf[i] + 1
			break
		}
	}
	if i == -1 {
		buf1 := bytes.NewBufferString("1")
		buf1.Write(buf)
		return buf1.String()
	}
	return string(buf)
}

func Dec(n string) string {
	length := len(n)
	buf := []byte(n)
	for i := length - 1; i >= 0; i-- {
		if buf[i] == '0' {
			buf[i] = '9'
		} else {
			buf[i] = buf[i] - 1
			break
		}
	}
	if buf[0] == '0' {
		buf = buf[1:]
	}
	return string(buf)
}

func AbSub(n1, n2 string) string {
	len1 := len(n1)
	len2 := len(n2)
	if len1 < len2 || (len1 == len2 && n1 < n2) {
		tmp := n1
		n1 = n2
		n2 = tmp
	}
	len1 = len(n1)
	len2 = len(n2)
	var remain = []byte(n1)
	var jiewei = false
	i, j := len1-1, len2-1
	for j >= 0 {
		if jiewei {
			remain[i] -= 1
		}
		//		remain[i] -= n2[j]
		if remain[i] < n2[j] {
			remain[i] += 10 - n2[j]
			jiewei = true
		} else {
			remain[i] -= n2[j]
			jiewei = false
		}
		remain[i] += '0'

		i--
		j--
	}
	//	println(n1, n2, string(remain))
	for i >= 0 && jiewei {
		remain[i] -= 1
		if remain[i] < '0' {
			remain[i] += 10
			i--
		} else {
			break
		}
	}
	return strings.TrimLeft(string(remain), "0")
}

//func biggernearest(bigger, middle, smaller string) bool {
//	o, _ := strconv.ParseInt(middle, 10, 0)
//	b, _ := strconv.ParseInt(bigger, 10, 0)
//	s, _ := strconv.ParseInt(smaller, 10, 0)
//	return b-o < o-s
//}

// odd
// if n[half]==9 && inc(n[:half]).width change
//    return  makePalindromic(inc(n[:half]))
// else
//		return
// even
func PalindromOfPalindrom(n string) (string, string) {
	half := len(n) / 2
	even := (len(n)%2 == 0)
	var smaller, bigger string
	pre_half_dec := Dec(n[:half])
	pre_half_inc := Inc(n[:half])
	if even {
		if len(pre_half_dec) == half {
			smaller = makePalindromic(pre_half_dec, "")
		} else {
			smaller = makePalindromic(pre_half_dec, "9")
		}

		if len(pre_half_inc) == half {
			bigger = makePalindromic(pre_half_inc, "")
		} else {
			bigger = makePalindromic(pre_half_inc[:half], "0")
		}
	} else {
		m := string(n[half])
		//		println(n, half)
		if m != "0" {
			smaller = makePalindromic(n[:half], string(n[half]-1))
		} else if len(pre_half_dec) == half {
			smaller = makePalindromic(pre_half_dec, "9")
		} else {
			smaller = makePalindromic(pre_half_dec+"9", "")
		}

		if m != "9" {
			bigger = makePalindromic(n[:half], string(n[half]+1))
		} else if len(pre_half_inc) == half {
			bigger = makePalindromic(pre_half_inc, "0")
		} else {
			bigger = makePalindromic(pre_half_inc, "")

		}
	}

	//	println("got", smaller, n, bigger)
	//	if biggernearest(bigger, n, smaller) {
	//		return bigger
	//	}

	//	return smaller
	return smaller, bigger
}

func smallest(arr []string) int {
	var minstr string
	var min int
	for i, s := range arr {
		if minstr == "" || len(minstr) > len(s) || (len(minstr) == len(s) && minstr > s) {
			minstr = s
			min = i
		}
	}
	return min
}
