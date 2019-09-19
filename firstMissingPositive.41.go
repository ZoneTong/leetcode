package main

func firstMissingPositive(nums []int) int {
	i, l := 0, len(nums)
	for i < l {
		if nums[i] > 0 && nums[i] <= l && nums[nums[i]-1] != nums[i] {
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		} else {
			i++
		}
	}
	// fmt.Println(nums)
	for i = 0; i < l; i++ { // 不用 range 可能节约了4ms
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return l + 1
}
