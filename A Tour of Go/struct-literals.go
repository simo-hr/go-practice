package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = Vertex{}
	p  = &Vertex{1, 2}
	v4 = Vertex{Y: 4}
	v5 = Vertex{X: 8, Y: 4}
)

func main() {
	fmt.Println((*p).X)
	fmt.Println(v1, p, v2, v3, v4, v5)
}
