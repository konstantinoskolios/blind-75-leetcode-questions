package main

import(
	"fmt"
)

func main() {
    matrix1 := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
    matrix2 := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}
	fmt.Println(spiralOrder(matrix1))
	fmt.Println(spiralOrder(matrix2))
}

func spiralOrder(matrix [][]int) []int {
    result := make([]int, 0)

    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return result
    }

    top, bottom, left, right := 0, len(matrix)-1, 0, len(matrix[0])-1

    for top <= bottom && left <= right {
        for i := left; i <= right; i++ {
            result = append(result, matrix[top][i])
        }
        top++

        for i := top; i <= bottom; i++ {
            result = append(result, matrix[i][right])
        }
        right--

        if top <= bottom {
            for i := right; i >= left; i-- {
                result = append(result, matrix[bottom][i])
            }
            bottom--
        }

        if left <= right {
            for i := bottom; i >= top; i-- {
                result = append(result, matrix[i][left])
            }
            left++
        }
    }

    return result
}