package main

import "fmt"

func main() {
	bd := newBoard()
	for i, nd := range bd.nodes {
		fmt.Printf("%d - x: %d y: %d\n", i, nd.x, nd.y)
	}
}
