package main

func shuffle(nums []int, n int) []int {
	numLength := len(nums)
	ans := make([]int, numLength)
	for i := range n{
		ans[i + i] = nums[i]
		ans[i + i + 1] = nums[n + i]
	}
	return ans
}