package main

type graph struct {
	nodes []*node
}

type node struct {
	x, y int
	//children []*node
}

type board struct {
	graph
}

func newBoard() *board {
	out := new(board)

	//square grid 7x7:
	w := 7
	h := 7
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			nd := new(node)
			nd.x = i * 1000 / w
			nd.y = j * 1000 / h
			out.nodes = append(out.nodes, nd)
		}
	}

	return out
}
