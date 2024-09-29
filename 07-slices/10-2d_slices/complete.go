package main

import (
	"fmt"
	"strings"
)

func createMatrix(rows, cols int) [][]int {
	matrix := [][]int{}
	for i := 0; i < rows; i++ {
		row := []int{}
		for j := 0; j < cols; j++ {
			row = append(row, i*j)
		}
		matrix = append(matrix, row)
	}
	return matrix
}

// dont edit below this line

func test(rows, cols int) {
	fmt.Printf("Creating %v x %v matrix...\n", rows, cols)
	matrix := createMatrix(rows, cols)
	for i := 0; i < len(matrix); i++ {
		fmt.Println(matrix[i])
	}
	fmt.Println("===== END REPORT =====")
}

func test2(rows, cols, padding int, padChar rune) {
	fmt.Printf("Creating %v x %v matrix...\n", rows, cols)
	matrix := createMatrix(rows, cols)

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			value := fmt.Sprintf("%d", matrix[i][j])
			if len(value) < padding {
				value = strings.Repeat(string(padChar), padding-len(value)) + value
			}
			fmt.Printf("%s ", value)
		}
		fmt.Println()
	}
	fmt.Println("===== END REPORT =====")
}

func main() {
	test(3, 3)
	test(5, 5)
	test(10, 10)
	test(15, 15)
	test2(3, 10, 3, ' ')
}
