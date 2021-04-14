package main

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type graph struct {
	nodes []*Node
}

type Node struct {
	x, y     float64
	children []*Node
	step     int
	start    bool
	end      bool
	block    bool
	path     bool
}

type Board struct {
	graph
	startNode *Node
	endNode   *Node
}

func newBoard() *Board {
	out := new(Board)

	//square grid:
	w := 12
	h := 9

	nodes := make([][]Node, w)
	for i := range nodes {
		nodes[i] = make([]Node, h)
	}

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			nodes[i][j].y = float64((j + 1) * 480 / (h + 1))
			nodes[i][j].x = float64((i + 1) * 640 / (w + 1))
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

	nodes[2][3].start = true
	nodes[9][7].end = true

	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			out.nodes = append(out.nodes, &nodes[i][j])

			if nodes[i][j].start {
				out.startNode = &nodes[i][j]
			}
			if nodes[i][j].end {
				out.endNode = &nodes[i][j]
			}
		}
	}

	return out
}

func (b *Board) draw(screen *ebiten.Image) {
	for _, node := range b.nodes {
		node.drawConnections(screen)
	}
	for _, node := range b.nodes {
		node.draw(screen)
	}
}

func (n *Node) drawConnections(screen *ebiten.Image) {
	for _, n2 := range n.children {
		ebitenutil.DrawLine(screen, n.x, n.y, n2.x, n2.y, color.RGBA{192, 192, 192, 255})
	}
}

func (n *Node) draw(screen *ebiten.Image) {
	const r = 16

	var col color.RGBA
	switch {
	case n.path:
		col = color.RGBA{255, 255, 0, 255}
	case n.block:
		col = color.RGBA{127, 127, 127, 255}
	case n.start:
		col = color.RGBA{0, 255, 255, 255}
	case n.end:
		col = color.RGBA{0, 0, 255, 255}
	default:
		col = color.RGBA{96, 64, 96, 255}
	}

	ebitenutil.DrawRect(screen, n.x-r, n.y-r, r*2, r*2, col)
	ebitenutil.DebugPrintAt(screen, strconv.Itoa(n.step), int(n.x)-r, int(n.y)-r)
}

func (b *Board) bfs() {
	queue := []*Node{b.startNode}
	queue[0].step = 1

	var foundNode *Node
	for foundNode == nil && len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		for _, ch := range node.children {
			if !ch.block && ch.step == 0 {
				ch.step = node.step + 1

				if ch == b.endNode {
					foundNode = node
					break
				}
				queue = append(queue, ch)
			}
		}
	}

	if foundNode != nil {
		node := foundNode
		for node != b.startNode {
			node.path = true
			for _, ch := range node.children {
				if ch.step == node.step-1 {
					node = ch
					break
				}
			}
		}
	}
}

func (n *Node) containsPoint(x, y int) bool {
	dx := (x - int(n.x)) * (x - int(n.x))
	dy := (y - int(n.y)) * (y - int(n.y))
	return dx*dx+dy*dy <= 128*128
}

func (b *Board) click() {
	for _, node := range b.nodes {
		x, y := ebiten.CursorPosition()
		if node.containsPoint(x, y) {
			node.block = !node.block
		}
	}
}

func (b *Board) clearNodes() {
	for _, n := range b.nodes {
		n.path = false
		n.step = 0
	}
}
