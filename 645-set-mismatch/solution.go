package main

func findErrorNums(nums []int) []int {
	n := len(nums)
	count := make([]int, n+1)

	for _, num := range nums {
		count[num]++
	}

	missing := 0
	duplicate := 0
	for i:= 1; i <= n; i++{
		if missing != 0 && duplicate !=0{
			return []int{duplicate, missing}
		}
		if count[i] == 2{
			duplicate = i
		}
		if count[i] == 0{
			missing = i
		}
	}

	return []int{duplicate, missing}
}
