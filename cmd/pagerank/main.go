package main

import (
	"fmt"

	"github.com/alixaxel/pagerank"
)

func main() {
	graph := pagerank.NewGraph()

	graph.Link(1, 2, 1.0)
	graph.Link(1, 3, 2.0)
	graph.Link(2, 3, 3.0)
	graph.Link(2, 4, 4.0)
	graph.Link(3, 1, 5.0)

	graph.Rank(0.85, 0.000001, func(node uint32, rank float64) {
		fmt.Println("Node", node, "has a rank of", rank)
	})
}
