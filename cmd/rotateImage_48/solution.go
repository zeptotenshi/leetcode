func rotate(matrix [][]int)  {
	j := len(matrix) - 1

	for i := 0; i < len(matrix)/2; i++ {
		for x := 0 ; x < j; x++ {
			// top -> right
			t1 := matrix[i+x][j]
			matrix[i+x][j] = matrix[i][i+x]

			// right -> bottom
			t2 := matrix[j][j-x]
			matrix[j][j-x] = t1

			// bottom -> left
			t1 = matrix[j-x][i]
			matrix[j-x][i] = t2

			// left -> top
			matrix[i][i+x] = t1
		}

		j-=2
	}
}

// Input: matrix = [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
// Output: [[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]