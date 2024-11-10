package main

import "fmt"

// func one (x *int) {
// 	*x  = *x +1
// }
// func main() {
// 	 var n int = 100
// 	 one(&n)
// 	 fmt.Println(n)
// 	//  fmt.Println(n)
// 	//  fmt.Println(&n)
// 	//  fmt.Println(&n)
// 	//  var p *int = &n
// 	//  fmt.Println(p)
// 	//  fmt.Println(*p)
// }

// func main() {
// 	/*
// 	var p *int = new(int)
// 	fmt.Println(p)
// 	fmt.Println(*p)
// 	*p++
// 	fmt.Println(*p)

// 	var p2 *int
// 	fmt.Println(p2)
// 	// *p2++ // error
// 	*/

// 	// makeとnewの違いは、ポインタを返すかどうか
// 	s := make([]int, 0)
// 	fmt.Printf("%T\n", s)

// 	m := make(map[string]int)
// 	fmt.Printf("%T\n", m)

// 	ch := make(chan int)
// 	fmt.Printf("%T\n", ch)

// 	var p *int = new(int)
// 	fmt.Printf("%T\n", p)

// 	var st = new(struct{})
// 	fmt.Printf("%T\n", st)
// }

type Vertex struct {
	X, Y int
	S    string
}

func ChangeVertex(v Vertex) {
	v.X = 1000
}

func ChangeVertex2(v *Vertex) {
	v.X = 1000
	// (*v).X = 1000
}



func main() {
	v := Vertex{1, 2, "test"}
	ChangeVertex(v)
	fmt.Println(v)

	v2 := Vertex{1, 2, "test"}
	ChangeVertex2(&v2)
	fmt.Println(v2)

	var i int = 100
	var j int = 200
	var p1 *int
	var p2 *int
	p1 = &i
	p2 = &j
	i = *p1 + *p2
	p2 = p1
	fmt.Println(*p2)
	fmt.Println(*p1)
	fmt.Println(i)
	j = *p2 + i
	fmt.Println(j)
	/*
		y := Vertex{X: 1, Y: 2}
		fmt.Println(y)
		fmt.Println(y.X, y.Y)

		v2 := Vertex{X: 1} // Yには0が入る
		fmt.Println(v2)

		v3 := Vertex{1, 2, "test"}
		fmt.Println(v3)

		v4 := Vertex{}
		fmt.Printf("%T %v\n",v4,v4)

		var v5 Vertex
		fmt.Printf("%T %v\n",v5,v5)

		v6 := new(Vertex)
		fmt.Printf("%T %v\n",v6,v6)

		v7 := &Vertex{}
		fmt.Printf("%T %v\n",v7,v7)
	*/
}
