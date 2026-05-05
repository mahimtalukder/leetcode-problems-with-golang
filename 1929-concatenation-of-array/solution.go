package main

func getConcatenation(nums []int) []int {
	n := len(nums)
	ans := make([]int, 2*n)

	k := 0
	for i := range ans{
		if i == n{
			k = 0
		}
		ans[i] = nums[k]
		k++
	}

	return ans
}