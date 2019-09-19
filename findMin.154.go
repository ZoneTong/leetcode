package main

// 解法2
// func findMin(nums []int) int {
// 	l, r := 0, len(nums)-1
// 	var mid int
// 	for l < r {
// 		mid = (l + r) / 2
// 		if nums[mid] > nums[r] {
// 			l = mid + 1
// 		} else if nums[mid] < nums[r] {
// 			r = mid
// 		} else {
// 			r--
// 		}
// 	}
// 	return nums[l]
// }

// 解法1: 分治,效率内存比解法2还快,易懂
func findMin(nums []int) int {
	if len(nums) == 1 || nums[0] < nums[len(nums)-1] {
		return nums[0]
	}

	mid := len(nums) >> 1
	min1 := findMin(nums[:mid])
	min2 := findMin(nums[mid:])
	if min1 > min2 {
		return min2
	}
	return min1
}
