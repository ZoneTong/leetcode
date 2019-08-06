package main

//	. "github.com/ZoneTong/leetcode/common"

//返回 A 的最短的非空连续子数组的长度，该子数组的和至少为 K 。
//如果没有和至少为 K 的非空子数组，返回 -1 。
func shortestSubarray(A []int, K int) int {
	min := len(A) + 1
	sum := 0
	begin := 0
	for i := 0; i < len(A); i++ {
		if A[i] >= K {
			return 1
		}

		sum += A[i]
		if sum < 1 {
			sum = 0
			begin = i + 1
			continue
		}

		// 合并成正数
		for j := i - 1; A[j+1] < 0; j-- {
			A[j] = A[j] + A[j+1]
			A[j+1] = 0
		}

		if sum < K {
			continue
		}

		for sum-A[begin] >= K || A[begin] <= 0 {
			sum -= A[begin]
			begin++
		}

		new_res := i - begin + 1
		if new_res < min {
			min = new_res
		}

	}

	if min == len(A)+1 {
		return -1
	}
	return min
}
