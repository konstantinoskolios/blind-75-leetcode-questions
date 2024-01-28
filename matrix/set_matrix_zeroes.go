package main

import (
	"fmt"
)

func main() {
    matrix1 := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
    setZeroes(matrix1)
    fmt.Println(matrix1) 

    matrix2 := [][]int{{0, 1, 2, 0}, {3, 4, 5, 2}, {1, 3, 1, 5}}
    setZeroes(matrix2)
    fmt.Println(matrix2) 
}

func setZeroes(matrix [][]int) {
    m := len(matrix)
    n := len(matrix[0])
    
    zeroRows := make([]bool, m)
    zeroCols := make([]bool, n)
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if matrix[i][j] == 0 {
                zeroRows[i] = true
                zeroCols[j] = true
            }
        }
    }
    
    for i := 0; i < m; i++ {
        if zeroRows[i] {
            for j := 0; j < n; j++ {
                matrix[i][j] = 0
            }
        }
    }
    
    for j := 0; j < n; j++ {
        if zeroCols[j] {
            for i := 0; i < m; i++ {
                matrix[i][j] = 0
            }
        }
    }
}


