package main

import (
	"testing"
)

func ASSERTION(t *testing.T, func1, val interface{}) {
	if func1 != val {
		t.Fatalf("need %v, but got %v", val, func1)
	}
}

func aTestShortestSubarray(t *testing.T) {
	ASSERTION(t, shortestSubarray([]int{1}, 1), 1)
	ASSERTION(t, shortestSubarray([]int{1, 2}, 4), -1)
	ASSERTION(t, shortestSubarray([]int{2, -1, 2}, 3), 3)
	//	ASSERTION(t, shortestSubarray([]int{484, 743, -42, 110, -294, -324, 610, 25, -19, 638, 382, 444, -57, -29, 19, 780, 914, 641, 558, -354, -410, 773, -175, -322, 564, -74, 699}, 1000), 3)
	ASSERTION(t, shortestSubarray([]int{56, -21, 56, 35, -9}, 61), 2)
}

func TestNearestPalindromic(t *testing.T) {
	ASSERTION(t, nearestPalindromic("123"), "121")
	ASSERTION(t, nearestPalindromic("1"), "0")
	ASSERTION(t, nearestPalindromic("9"), "8")
	ASSERTION(t, nearestPalindromic("10"), "9")
	ASSERTION(t, nearestPalindromic("11"), "9")
	ASSERTION(t, nearestPalindromic("22"), "11")
	ASSERTION(t, nearestPalindromic("99"), "101") //L 9&&change
	ASSERTION(t, nearestPalindromic("101"), "99")
	ASSERTION(t, nearestPalindromic("191"), "181")
	ASSERTION(t, nearestPalindromic("202"), "212")  //L 191 odd&&0
	ASSERTION(t, nearestPalindromic("909"), "919")  //L 898 odd&&0
	ASSERTION(t, nearestPalindromic("999"), "1001") //L odd&&9&&change
	ASSERTION(t, nearestPalindromic("1001"), "999")
	ASSERTION(t, nearestPalindromic("1991"), "2002")  //L 9&&change
	ASSERTION(t, nearestPalindromic("2002"), "1991")  // 2112
	ASSERTION(t, nearestPalindromic("9009"), "8998")  // 9119
	ASSERTION(t, nearestPalindromic("9999"), "10001") //L 9&&change
	ASSERTION(t, nearestPalindromic("10001"), "9999")
	ASSERTION(t, nearestPalindromic("20002"), "19991")
	ASSERTION(t, nearestPalindromic("90009"), "89998")     // 90109
	ASSERTION(t, nearestPalindromic("2400042"), "2399932") // 2401042
}

//func strparse(n string) int {
//	i, _ := strconv.ParseInt(n, 10, 0)
//	return int(i)
//}

func aTestStringIncDec(t *testing.T) {
	t.Log(Inc("90"))
	t.Log(Inc("99"))
	t.Log(Inc("199"))
	t.Log(Dec("100"))
	t.Log(Dec("110"))
	t.Log(Dec("200"))
	t.Log(Dec("201"))
	t.Log(Dec("1"))
	t.Log(AbSub("202", "199"))
	t.Log(AbSub("202", "206"))
}
