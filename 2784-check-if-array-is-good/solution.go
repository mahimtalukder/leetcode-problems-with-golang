package main

func isGood(nums []int) bool {
    maxNum := 0

	for _, num := range nums{
		if maxNum < num {
			maxNum = num
		}
	}

	if len(nums) != maxNum + 1{
		return false
	}

	counts := make([]int, maxNum+1)

	for _, num := range nums{
		counts[num]++ 
	}

	for i := 1; i<maxNum+1; i++{
		if i != maxNum && counts[i] > 1{
			return false
		}else if i==maxNum && counts[i] != 2{
			return false
		}

	}

	return true
}