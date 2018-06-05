package main

import "fmt"

func main() {
	n := 4
	matrix := [][]string{
		[]string{"a", "b", "c", "d"},
		[]string{"e", "f", "g", "h"},
		[]string{"i", "j", "k", "l"},
		[]string{"m", "n", "o", "p"},
	}
	fmt.Println("before rotation:")
	printMatrix(matrix, n)
	fmt.Println()
	fmt.Println("after rotation:")
	rotation(matrix, n)
	printMatrix(matrix, n)

}

//Given an image represented by an NxN matrix,
// where each pixel in the image is 4 bytes,
// write a method to rotate the image by 90 degrees. Can you do this in place?

//[1][2][3][4]
//[5][6][7][8]
//[9][0][1][2]
//[3][4][5][6]
//Becomes:
//[3][9][5][1]
//[4][0][6][2]
//[5][1][7][3]
//[6][2][8][4]

//The rotation can be performed in layers, where you perform a cyclic swap on the edges on each layer.
// In the first for loop, we rotate the first layer (outermost edges).
// We rotate the edges by doing a four-way swap first on the corners, then on the element clockwise from the edges, then on the element three steps away.
//Once the exterior elements are rotated, we then rotate the interior region’s edges.
//https://stackoverflow.com/questions/42519/how-do-you-rotate-a-two-dimensional-array
func rotation(matrix [][]string, n int) {
	// layer是rotatable layer number，其数值为n/2, for example, a 7*7 metrics has 3 rotatable layers
	//. . . . . . .
	//. x x x x x .
	//. x O O O x .
	//. x O - O x .
	//. x O O O x .
	//. x x x x x .
	//. . . . . . .
	//loop each layer
	for layer := 0; layer < n/2; layer++ {
		//we need to increment the the row and decrement the column positions when moving inwards to the next layer.
		//The variables first and last identify the index position of the first and last rows and columns.
		first := layer
		last := n - 1 - layer
		//Now we need a way of navigating within a layer so we can move elements around that layer.
		//loop each element in the layer
		for element := first; element < last; element++ {
			//. . . . .
			//. x x x .
			//. x O x .
			//. x x x .
			//. . . . .
			offset := element - first
			top := matrix[first][element] // save top

			// 想象绕着layer，沿顺时针遍历element，
			// 从top开始向右，top: 行 = first, 列逐渐增加 = element
			// 接下俩是右边向下，行逐渐增加 =element, 列 = last
			// 然后是底部向左，行=last，列组建减少=last-offset
			// 最后是左侧向上,行逐渐减少=last-offset，列=first
			//top = matrix[first][element], Moving forwards along the top row requires the column index to be incremented.
			//right_side = matrix[element][last], Moving down the right side requires the row index to be incremented.
			//bottom = matrix[last][last-offset], Moving backwards along the bottom requires the column index to be decremented.
			//left_side = matrix[last-offset][first], Moving up the left side requires the row index to be decremented.

			//matrix[first][element] = left_side
			//matrix[element][last] = top
			//matrix[last][last-offset] = right_side
			//matrix[last-offset][first] = bottom

			// left -> top
			matrix[first][element] = matrix[last-offset][first]
			// bottom -> left
			matrix[last-offset][first] = matrix[last][last-offset]
			// right -> bottom
			matrix[last][last-offset] = matrix[element][last]
			// top -> right
			matrix[element][last] = top // right <- saved top
		}
	}
}

func printMatrix(matrix [][]string, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Print("(", i, ",", j, ")is:", matrix[i][j], ", ")

		}
		fmt.Println()
	}
}
