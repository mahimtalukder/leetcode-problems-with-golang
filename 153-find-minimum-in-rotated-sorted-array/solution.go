package main

func findMin(nums []int) int {
	min := 5001

	for _, num := range nums{
		if min > num{
			min = num
		}
	}
    return min
}