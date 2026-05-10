package main

func rotateGrid(grid [][]int, k int) [][]int {
	m := len(grid)
	n := len(grid[0])

	ans := make([][]int, m)
	for i := range ans {
		ans[i] = make([]int, n)
		copy(ans[i], grid[i])
	}

	layers := min(m, n) / 2

	for layer := range layers {
		indexes := getLayerIndexes(m, n, layer)
		p := len(indexes)
		step := k % p

		for i := range p {
			from := indexes[i]
			to := indexes[(i+step)%p]

			ans[to[0]][to[1]] = grid[from[0]][from[1]]
		}
	}

	return ans
}

func getLayerIndexes(m, n, layer int) [][2]int {
	top := layer
	left := layer
	bottom := m - 1 - layer
	right := n - 1 - layer

	indexes := make([][2]int, 0)

	// down left column
	for r := top; r <= bottom; r++ {
		indexes = append(indexes, [2]int{r, left})
	}

	// right bottom row
	for c := left + 1; c <= right; c++ {
		indexes = append(indexes, [2]int{bottom, c})
	}

	// up right column
	for r := bottom - 1; r >= top; r-- {
		indexes = append(indexes, [2]int{r, right})
	}

	// left top row
	for c := right - 1; c > left; c-- {
		indexes = append(indexes, [2]int{top, c})
	}

	return indexes
}