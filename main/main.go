package main

import (
	"fmt"
	"github.com/leoxiang66/go-graph/galg/centrality"
)

func main() {
    // Define a 3x3 matrix
    matrix := [][]int{
        {0, 1, 0},
        {0, 0, 1},
        {0, 1, 0},
    }

	result := centrality.Article_Rank(matrix,100)

    // Print the entire matrix
    fmt.Println("Matrix:")
    for _, row := range matrix {
        fmt.Println(row)
    }

    fmt.Println("Result:")
    for _, row := range result {
        fmt.Println(row)
    }

    // Access specific elements
    fmt.Println("\nAccessing specific elements:")
    fmt.Printf("Element at (0, 1): %d\n", matrix[0][1]) // 2
    fmt.Printf("Element at (2, 2): %d\n", matrix[2][2]) // 9

    // Modify an element
    matrix[1][1] = 42
    fmt.Println("\nMatrix after modification:")
    for _, row := range matrix {
        fmt.Println(row)
    }
}