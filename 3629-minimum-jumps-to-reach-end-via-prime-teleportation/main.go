package main

import "fmt"

func main() {
	nums := []int{7, 17, 19, 23, 39, 11, 29, 31, 13, 14, 47, 53, 59, 61, 55}

	fmt.Println(nums)
	fmt.Print(minJumps(nums))
}