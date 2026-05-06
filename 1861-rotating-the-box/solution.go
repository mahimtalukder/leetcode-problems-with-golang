package main

func rotateTheBox(boxGrid [][]byte) [][]byte {
	row := len(boxGrid)
	col := len(boxGrid[0])

	// Step 1: Apply gravity to each column
	for i := 0; i < row; i++ {
		// Start from the bottom of each column to simulate gravity
		freeSpot := col - 1 // The position where the stone should fall

		// Traverse the column from the bottom
		for j := col - 1; j >= 0; j-- {
			if boxGrid[i][j] == '*' {
				freeSpot = j - 1 // An obstacle blocks the stones from falling further
			} else if boxGrid[i][j] == '#' {
				boxGrid[i][freeSpot] = '#' // Place the stone at the free spot
				if freeSpot != j { // If the stone moved, clear the original spot
					boxGrid[i][j] = '.'
				}
				freeSpot-- // Move to the next free spot
			}
		}
	}

	// Step 2: Rotate the matrix 90 degrees clockwise
	rotatedBoxGrid := make([][]byte, col)
	for i := 0; i < col; i++ {
		rotatedBoxGrid[i] = make([]byte, row)
		for j := 0; j < row; j++ {
			rotatedBoxGrid[i][row-1-j] = boxGrid[j][i] // Swap rows and columns for rotation
		}
	}

	return rotatedBoxGrid
}