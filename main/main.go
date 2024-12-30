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

	

    // Print the entire matrix
    fmt.Println("Adjacency Matrix:")
    for _, row := range matrix {
        fmt.Println(row)
    }

    fmt.Println("\n\nResult of Article Rank:")
    result := centrality.Article_Rank(matrix,10,0.8)
    for _, row := range result {
        fmt.Println(row)
    }


}