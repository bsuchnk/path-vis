package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DrawLine(screen, 10, 10, 100, 100, color.RGBA{123, 43, 0, 255})
}

func (g *Game) Layout(oustideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("title")
	game := &Game{}
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}

	// bd := newBoard()
	// for i, nd := range bd.nodes {
	// 	fmt.Printf("%d - x: %d y: %d\n", i, nd.x, nd.y)
	// }
}
