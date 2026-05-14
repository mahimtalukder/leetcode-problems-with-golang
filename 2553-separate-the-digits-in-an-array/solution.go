package main

import "strconv"

func separateDigits(nums []int) []int {
	ans := []int{}
	for _, num := range nums {
		if num > 9 {
			if num < 100 {
				ans = append(ans, num/10, num%10)
			} else {
				n := len(strconv.Itoa(num))
				convertedInt := make([]int, n)
				for i := n-1; i >= 1; i--{
					convertedInt[i] = num%10
					num = num/10
				}
				convertedInt[0] = num
				ans = append(ans, convertedInt...)
			}
		} else {
			ans = append(ans, num)
		}
	}
	return ans
}
