package main

import "fmt"

func main() {
	//nums := []int{11,18,11}
	//nums := []int{12, 12, 14, 11, 11, 12 , 10, 40, 45, 10, 10, 125, 50, 15}
	//nums := []int{2,1,3}
	nums := []int{8,5,20,2,21}
	//nums := []int{12, 12, 14, 11, 11, 12 , 10, 40, 45, 10, 10, 125, 50, 15}
	//nums := []int{56,8,56,31,68,15,87,75,63}

	fmt.Println(nums)
	fmt.Print(maxValue(nums))
}
