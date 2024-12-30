package centrality

import "github.com/leoxiang66/go-graph/utils"

func Article_Rank(graph [][]int, iteration int, damping float32) []float32{
	ret := make([]float32, 1) // ranks
	num_nodes := len(graph)

	in,in_degs := utils.Get_In_Nodes(graph)
	out,out_degs := utils.Get_Out_Nodes(graph)

	for i := 0; i < iteration; i++ {
		// for every node, compute new rank value
		for j := 0; j < num_nodes; j++ {
			new_rank := 1 - damping

		}
	}


	return ret
}