package main

import "fmt"

func main() {
	tasks := [][]int{
		{1, 3},
		{2, 4},
		{10, 11},
		{10, 12},
		{8, 9},
	}

	fmt.Println(minimumEffort(tasks))
}
