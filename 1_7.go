package main

import (
	"fmt"
	"reflect"
)

func main() {
	type TestSetZeros struct {
		givenMatrix    [][]int
		expectedMatrix [][]int
	}

	cases := []TestSetZeros{
		{
			[][]int{},
			[][]int{},
		},
		{
			[][]int{
				{1, 1, 1, 0, 1}},
			[][]int{
				{0, 0, 0, 0, 0}},
		},
		{
			[][]int{
				{1, 1, 1, 1, 1}},
			[][]int{
				{1, 1, 1, 1, 1}},
		},
		{
			[][]int{
				{1},
				{0},
				{1},
				{1}},
			[][]int{
				{0},
				{0},
				{0},
				{0}},
		},
		{
			[][]int{
				{1},
				{1},
				{1},
				{1}},
			[][]int{
				{1},
				{1},
				{1},
				{1}},
		},
		{
			[][]int{
				{1, 1, 1, 0, 1},
				{0, 1, 1, 1, 1},
				{1, 1, 1, 1, 1},
				{1, 0, 1, 1, 1}},
			[][]int{
				{0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0},
				{0, 0, 1, 0, 1},
				{0, 0, 0, 0, 0}},
		},
	}

	for _, c := range cases {
		setZeros(c.givenMatrix)
		fmt.Println(reflect.DeepEqual(c.expectedMatrix, c.givenMatrix))
	}

}

//Write an algorithm such that if an element in an MxN matrix is 0, its entire row and column is set to 0.
func setZeros(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	zeroRaws := make([]bool, len(matrix))
	zeroColumns := make([]bool, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == 0 {
				zeroRaws[i] = true
				zeroColumns[j] = true
			}
		}
	}

	// Set matrix[i][j] to 0 if either row i or column j has a 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if zeroRaws[i] == true || zeroColumns[j] == true {
				matrix[i][j] = 0
			}
		}
	}

}

//func printRectangleMatrix(matrix [][]int) {
//	for i := 0; i < len(matrix); i++ {
//		for j := 0; j < len(matrix[0]); j++ {
//			fmt.Print("(", i, ",", j, ")is:", matrix[i][j], ", ")
//
//		}
//		fmt.Println()
//	}
//	fmt.Println()
//}
