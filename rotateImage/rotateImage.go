package rotateImage

import (
	"fmt"
)

/* Solution Accepted
	Runtime: 0 ms
		faster than 100.00% of Go online submissions
	Memory Usage: 2.1 MB
		less than 89.71% of Go online submissions
*/

func rotate(matrix [][]int)  {
	// column index; marks max element of ring
	j := len(matrix) - 1
	// row index; marks first element of ring
	for i := 0; i < len(matrix)/2; i++ {
		fmt.Printf("\ni = %d, j = %d", i, j)
		for x := 0; x < j - i; x++ {
			fmt.Printf("\nx = %d", x)

			// rotate top edge -> right edge
			t1 := matrix[i+x][j]
			matrix[i+x][j] = matrix[i][i+x]
			fmt.Printf("\n\tcache <= matrix[%d][%d] = %d\n\tmatrix[%d][%d] <= matrix[%d][%d] = %d", i+x, j, t1, i+x, j, i, i+x, matrix[i+x][j])

			// rotate right edge -> bottom edge
			t2 := matrix[j][j-x]
			matrix[j][j-x] = t1
			fmt.Printf("\n\n\tcache <= matrix[%d][%d] = %d\n\tmatrix[%d][%d] <= matrix[%d][%d] = %d", j, j-x, t2, j, j-x, i+x, j, matrix[j][j-x])

			// rotate bottom edge -> left edge
			t1 = matrix[j-x][i]
			matrix[j-x][i] = t2
			fmt.Printf("\n\n\tcache <= matrix[%d][%d] = %d\n\tmatrix[%d][%d] <= matrix[%d][%d] = %d", j-x, i, t1, j-x, i, j, j-x, matrix[j-x][i])

			// rotate left edge -> top edge
			matrix[i][i+x] = t1
			fmt.Printf("\n\tmatrix[%d][%d] = %d\n", i, i+x, matrix[i][i+x])
		}

		// decrement max column index
		j--
	}
}
