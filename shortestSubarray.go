package main

//	. "github.com/ZoneTong/leetcode/common"

//返回 A 的最短的非空连续子数组的长度，该子数组的和至少为 K 。
//如果没有和至少为 K 的非空子数组，返回 -1 。
func shortestSubarray(A []int, K int) int {
	min := len(A)
	sum := 0
	for i := 0; i < len(A); i++ {
		sum = 0
		for j := 0; j < min && j+i < len(A); j++ {
			//			if i == 2 {
			//				println("i=2", i, j, min)
			//			}

			sum += A[j+i]

			if sum >= K {
				min = j
				//				println("got", i, j, sum)
				break
			} else if sum < 1 {
				//				println(i, j, min)
				i = j + i // i++ will let it +1
				break
			}

		}
	}
	if min == len(A) {
		return -1
	}
	return min + 1
}
