package main

func rotate(matrix [][]int) {
	//transpose the matrix
	n := len(matrix)
	for i := range matrix {
		for j := i + 1; j < n; j++{
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	//reverse the matrix
	for i := range matrix {
		for j:=0; j<n/2; j++{
			k := n - (j + 1)
			matrix[i][j], matrix[i][k] = matrix[i][k], matrix[i][j]
		}
	}
}
