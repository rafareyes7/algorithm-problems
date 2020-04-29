package main

import (
	"fmt"
	"reflect"
)

/*
Write a function that takes in an n x m two-dimensional array
(that can be square-shaped when n === m) and returns a one-dimensional
array of all the array's elements in spiral order.

Spiral order starts at the top left corner of the two-dimensional array,
goes to the right, and proceeds in a spiral pattern all the way until
every element has been visited.

example:

input: [
	[1, 2, 3, 4],
	[12, 13, 14, 5],
	[11, 16, 15, 6],
	[10, 9, 8, 7]
]

output: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16]
*/

func spiralTraverse(array [][]int) []int {
	if len(array) == 0 {
		return []int{}
	}
	result := []int{}

	// Set bounds
	startRow, endRow := 0, len(array)-1
	startCol, endCol := 0, len(array[startRow])-1

	for startRow <= endRow && startCol <= endCol {
		for col := startCol; col <= endCol; col++ {
			result = append(result, array[startRow][col])
		}

		for row := startRow + 1; row <= endRow; row++ {
			result = append(result, array[row][endCol])
		}

		for col := endCol - 1; col >= startCol; col-- {
			// If there is a single row at the middle of the matrix, at this
			// point we have already appended it to the result's array in the first loop.
			if startRow == endRow {
				break
			}
			result = append(result, array[endRow][col])
		}

		for row := endRow - 1; row > startCol; row-- {
			// If there is a single column in the middle of the matrix, at this
			// point we have already appended it to the result's array in the second loop.
			if startCol == endCol {
				break
			}
			result = append(result, array[row][startCol])
		}

		startRow++
		startCol++
		endRow--
		endCol--
	}

	return result
}

func main() {
	want := []int{1, 2, 3, 4}
	got := spiralTraverse([][]int{[]int{1, 2}, []int{4, 3}})
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v\n", want, got)
	}

	// Edge case: single row in th middle of the matrix.
	want = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	got = spiralTraverse([][]int{[]int{1, 2, 3, 4}, []int{10, 11, 12, 5}, []int{9, 8, 7, 6}})
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v\n", want, got)
	}

	// Edge case: single column in th middle of the matrix.
	want = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	got = spiralTraverse([][]int{[]int{1, 2, 3}, []int{12, 13, 4}, []int{11, 14, 5}, []int{10, 15, 6}, []int{9, 8, 7}})
	if !reflect.DeepEqual(want, got) {
		fmt.Printf("want: %v, got: %v\n", want, got)
	}

}
