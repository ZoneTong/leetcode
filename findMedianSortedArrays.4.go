// 4. 寻找两个有序数组的中位数
// https://leetcode-cn.com/problems/median-of-two-sorted-arrays/
package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// o(m+n)
	// h1, t1 := 0, len(nums1)-1
	// h2, t2 := 0, len(nums2)-1

	// for h1 <= t1 && h2 <= t2 {
	// 	if h1 == t1 && h2 == t2 {
	// 		return float64(nums1[h1]+nums2[h2]) / 2
	// 	}

	// 	if nums1[h1] < nums2[h2] {
	// 		h1++
	// 	} else {
	// 		h2++
	// 	}

	// 	if nums1[t1] < nums2[t2] {
	// 		t2--
	// 	} else {
	// 		t1--
	// 	}
	// }

	// var s int
	// var nums []int
	// if h1 <= t1 {
	// 	s = h1 + t1
	// 	nums = nums1
	// } else if h2 <= t2 {
	// 	s = h2 + t2
	// 	nums = nums2
	// }

	// return float64(nums[s/2]+nums[(s+1)/2]) / 2

	// 以下为O(log(m+n))解法
	l1 := len(nums1)
	l2 := len(nums2)
	if l1 > l2 {
		nums1, nums2 = nums2, nums1
		l1, l2 = l2, l1
	}

	// 奇数个数时将中位数划在左边
	iMin, iMax, half := 0, l1, (l1+l2+1)/2
	for iMin <= iMax {
		i := (iMin + iMax) / 2
		j := half - i
		if i > iMin && nums1[i-1] > nums2[j] {
			iMax = i - 1
		} else if i < iMax && nums1[i] < nums2[j-1] {
			iMin = i + 1
		} else {
			var left int
			if i == 0 {
				left = nums2[j-1]
			} else if j == 0 {
				left = nums1[i-1]
			} else if nums1[i-1] < nums2[j-1] {
				left = nums2[j-1]
			} else {
				left = nums1[i-1]
			}

			if (l1+l2)%2 == 1 {
				return float64(left)
			}

			var right int
			if i == l1 {
				right = nums2[j]
			} else if j == l2 {
				right = nums1[i]
			} else if nums1[i] < nums2[j] {
				right = nums1[i]
			} else {
				right = nums2[j]
			}

			return float64(left+right) / 2
		}

	}
	return 0
}
