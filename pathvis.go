package main

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	board Board
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.board.draw(screen)
	//ebitenutil.DrawLine(screen, 10, 10, 100, 100, color.RGBA{123, 43, 0, 255})
}

func (g *Game) Layout(oustideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("PathVis")
	game := &Game{}
	game.board = *newBoard()
	game.board.bfs()
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}

	// bd := newBoard()
	// for i, nd := range bd.nodes {
	// 	fmt.Printf("%d - x: %d y: %d\n", i, nd.x, nd.y)
	// }
}
