package main


func findMaxConsecutiveOnes(nums []int) int {
    maxOneCount := 0
	currentSequenceOneCount := 0

	for _, num := range nums{
		switch num {
		case 1:
			currentSequenceOneCount++
		case 0:
			if maxOneCount < currentSequenceOneCount{
				maxOneCount = currentSequenceOneCount
			}
			currentSequenceOneCount = 0
		}
	}

	if maxOneCount < currentSequenceOneCount{
		maxOneCount = currentSequenceOneCount
	}

	return maxOneCount
}