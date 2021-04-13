package main

type graph struct {
	nodes []*node
}

type node struct {
	x, y     int
	children []*node
}

type board struct {
	graph
}

func newBoard() *board {
	out := new(board)

	//square grid 7x7:
	w := 7
	h := 7

	nodes := make([][]node, w)
	for i := range nodes {
		nodes[i] = make([]node, h)
	}

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			nodes[i][j].y = j * 1000 / h
			nodes[i][j].x = i * 1000 / w
		}
	}

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			if i > 0 {
				nodes[i][j].children = append(nodes[i][j].children, &nodes[i-1][j])
			}
			if j > 0 {
				nodes[i][j].children = append(nodes[i][j].children, &nodes[i][j-1])
			}
			if i < w-1 {
				nodes[i][j].children = append(nodes[i][j].children, &nodes[i+1][j])
			}
			if j < h-1 {
				nodes[i][j].children = append(nodes[i][j].children, &nodes[i][j+1])
			}
		}
	}

	return out
}

func (b *board) draw() {

}
