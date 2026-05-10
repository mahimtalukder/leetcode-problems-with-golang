package main

import "fmt"

func main() {
	grid := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10, 11, 12},
		{13, 14, 15, 16},
	}

	k := 3

	printGrid(grid)
	fmt.Println("Output:")
	rotatedGrid := rotateGrid(grid, k)

	printGrid(rotatedGrid)
}

func printGrid(grid [][]int) {
	for _, row := range grid {
		for _, val := range row {
			fmt.Printf("%3d", val)
		}
		fmt.Println()
	}
}