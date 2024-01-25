package main

import "fmt"

func pacificAtlantic(heights [][]int) [][]int {
    if len(heights) == 0 || len(heights[0]) == 0 {
        return nil
    }

    m, n := len(heights), len(heights[0])
    pacificVisited := make([][]bool, m)
    atlanticVisited := make([][]bool, m)

    for i := range pacificVisited {
        pacificVisited[i] = make([]bool, n)
        atlanticVisited[i] = make([]bool, n)
    }

    result := [][]int{}

    // DFS for Pacific Ocean
    for i := 0; i < m; i++ {
        dfs(heights, i, 0, pacificVisited)
    }

    for j := 0; j < n; j++ {
        dfs(heights, 0, j, pacificVisited)
    }

    // DFS for Atlantic Ocean
    for i := 0; i < m; i++ {
        dfs(heights, i, n-1, atlanticVisited)
    }

    for j := 0; j < n; j++ {
        dfs(heights, m-1, j, atlanticVisited)
    }

    // Check for cells that can flow to both oceans
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if pacificVisited[i][j] && atlanticVisited[i][j] {
                result = append(result, []int{i, j})
            }
        }
    }

    return result
}

func dfs(heights [][]int, r, c int, visited [][]bool) {
    if visited[r][c] {
        return
    }

    visited[r][c] = true

    directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

    for _, dir := range directions {
        nr, nc := r+dir[0], c+dir[1]

        if nr >= 0 && nr < len(heights) && nc >= 0 && nc < len(heights[0]) && heights[nr][nc] >= heights[r][c] && !visited[nr][nc] {
            dfs(heights, nr, nc, visited)
        }
    }
}

func main() {
    // Example usage:
    heights := [][]int{
        {1, 2, 2, 3, 5},
        {3, 2, 3, 4, 4},
        {2, 4, 5, 3, 1},
        {6, 7, 1, 4, 5},
        {5, 1, 1, 2, 4},
    }

    result := pacificAtlantic(heights)

    fmt.Println("Output:", result)
}
