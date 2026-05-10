package main

func findDisappearedNumbers(nums []int) []int {
	n := len(nums)
	allNums := make([]int, n+1)
	totalDuplicateValue := 0

	for _, num := range nums{
		if allNums[num] > 0{
			totalDuplicateValue++
		}
		allNums[num]++
	}

	ans := make([]int, totalDuplicateValue)
	startInx := 0
	for i := 1; i<=n; i++{
		if allNums[i] == 0{
			ans[startInx] = i
			startInx++
		}
	}

	return ans

}
