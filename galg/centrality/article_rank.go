package centrality

import (
	"github.com/leoxiang66/go-graph/utils"
)

func Article_Rank(graph [][]int, iteration int, damping float32) []float32{
	last_step := make([]float32, len(graph))
	current_step := make([]float32, len(graph))
	num_nodes := len(graph)

	in,_ := utils.Get_In_Nodes(graph)
	_,out_degs := utils.Get_Out_Nodes(graph)
	avg_out_deg := utils.Sum(out_degs) / len(out_degs)

	for i := 0; i < iteration; i++ {
		// for every node, compute new rank value
		for j := 0; j < num_nodes; j++ {
			new_rank := 1 - damping
			for _, in_node := range in[j] {
				new_rank += damping * last_step[in_node] / (float32)(out_degs[in_node] + avg_out_deg)
			}
			current_step[j] = new_rank 
		}
		last_step = current_step
	}


	return utils.Normalize(current_step)
}