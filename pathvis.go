package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Game struct {
	board Board
}

func (g *Game) Update() error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		g.board.click()
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		g.board.clearNodes()
		g.board.bfs()
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.board.draw(screen)
}

func (g *Game) Layout(oustideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("PathVis")
	game := &Game{}
	game.board = *newBoard()
	//game.board.bfs()
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
