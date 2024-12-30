package utils


// Normalize function normalizes an array of float64 values to the range [0, 1].
func Normalize[T int | float64 | float32](arr []T) []T {
	if len(arr) == 0 {
		return arr // Return empty array if input is empty
	}

	sum := Sum(arr)

	// Normalize the array
	normalized := make([]T, len(arr))
	for i, value := range arr {
		normalized[i] = value / sum 
	}

	return normalized
}

// Sum function that works with any numeric type
func Sum[T int | float64 | float32](arr []T) T {
	var sum T
	for _, value := range arr {
		sum += value
	}
	return sum
}

// Get_In_Nodes returns a 2D slice of incoming nodes and an array of in-degrees.
func Get_In_Nodes(graph [][]int) ([][]int, []int) {
	ret := make([][]int, len(graph))
	inDeg := make([]int, len(graph))

	for i := 0; i < len(graph); i++ {
		ret[i] = make([]int, 0) // Initialize the inner slice

		for j := 0; j < len(graph); j++ {
			if graph[j][i] == 1 {
				ret[i] = append(ret[i], j) // Append incoming node
				inDeg[i]++                  // Increment in-degree for node i
			}
		}
	}
	return ret, inDeg
}

// Get_Out_Nodes returns a 2D slice of outgoing nodes and an array of out-degrees.
func Get_Out_Nodes(graph [][]int) ([][]int, []int) {
	ret := make([][]int, len(graph))
	outDeg := make([]int, len(graph))

	for i := 0; i < len(graph); i++ {
		ret[i] = make([]int, 0) // Initialize the inner slice

		for j := 0; j < len(graph); j++ {
			if graph[i][j] == 1 {
				ret[i] = append(ret[i], j) // Append outgoing node
				outDeg[i]++                 // Increment out-degree for node i
			}
		}
	}
	return ret, outDeg
}