package main

func smallerNumbersThanCurrent(nums []int) []int {
	n := len(nums)
	maxValue := 0
	countArr := make([]int, 101)

	for _, num := range nums {
		if num > maxValue {
			maxValue = num
		}
		countArr[num]++
	}

	for i := maxValue; i >= 0; i-- {
		if(countArr[i] != 0){
			n = n - countArr[i]
			if n <= 0{
				countArr[i] = 0
			}else{
				countArr[i] = n
				n = countArr[i]
			}
		}
	}
	
	ans := make([]int, len(nums))
	for i, num := range nums{
		ans[i] = countArr[num]
	}

	return ans
}
