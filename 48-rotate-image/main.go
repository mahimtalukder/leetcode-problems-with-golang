package main

import "fmt"

func main() {
	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("matrix: ")
	for _, row := range matrix {
		for _, value := range row {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}

	fmt.Println("rotate matrix: ")
	rotate(matrix)
	for _, row := range matrix {
		for _, value := range row {
			fmt.Printf("%d ", value)
		}
		fmt.Println()
	}
}