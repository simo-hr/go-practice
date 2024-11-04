package main

import "fmt"

// import (
// 	"fmt"
// 	"os/user"
// 	"time"
// )

// func main() {
// 	fmt.Println("Hello, World!",time.Now())
// 	fmt.Println(user.Current())
// }

// import "fmt"
// func main() {
// 	var (i int = 1
// 	 f64 float64 = 1.2
// 	 s string = "test"
// 	 t,f bool = true,false)
// 	fmt.Println(i,f64,s,t,f)
// 	// var i int = 1
// 	// var f64 float64 = 1.2
// 	// var s string = "test"
// 	// var t,f bool = true,false
// 	// fmt.Println(i,f64,s,t,f)

// 	xi := 1
// 	xi = 2
// 	xf64 := 1.2
// 	var xf32 float32 = 1.2
// 	xs := "test"
// 	xt,xf := true,false
// 	fmt.Println(xi,xf64,xs,xt,xf)
// 	fmt.Printf("%T\n",xf64)
// 	fmt.Printf("%T\n",xf32)
// 	fmt.Printf("%T\n",xi)
// }

// // 定数はグローバルで宣言するのが一般的
// const Pi  = 3.14
// const (
// 	Username = "test_user"
// 	Password = "test_pass"
// )
// // var big int = 9223372036854775807 + 1
// const big = 9223372036854775807 + 1
// func main() {
// 	fmt.Println(Pi,Username,Password)
// 	// Pi = 3
// 	fmt.Println(big - 1)
// }

// func main() {
/*
	var (
		u8 uint8      = 255
		i8 int8       = 127
		f32 float32   = 0.2
		c64 complex64 = -5 + 12i

	)
	fmt.Println(u8,i8,f32,c64)
	fmt.Printf("type=%T value=%v",u8,u8)
*/

/*
	x := 1 + 1
	fmt.Println(x)
	fmt.Println("1 + 1 =", 1+1)
	fmt.Println("10 - 1 =", 10-1)
	fmt.Println("10 / 2 =", 10/2)
	fmt.Println("10 / 3 =", 10/3)
	fmt.Println("10.0 / 3 =", 10.0/3)
	fmt.Println("10 / 3.0 =", 10/3.0)
	fmt.Println("10 % 2 =", 10%2)
	fmt.Println("10 % 3 =", 10%3)
*/

// }

// func main() {
// 	fmt.Println("Hello, World!")
// 	fmt.Println("Hello" + "World")
// 	fmt.Println("Hello, World!"[0])         // 72（ASCIIコード）
// 	fmt.Println(string("Hello, World!"[0])) // 72（ASCIIコード）

// 	var s string = "Hello, World!"
// 	// s[0] = "X" // エラー
// 	s = strings.Replace(s, "H", "X", 1)
// 	fmt.Println(s)
// 	fmt.Println(strings.Contains(s, "World"))
// 	fmt.Println("Test\nTest")
// 	fmt.Println(`Test
// 	           Test
// Test`)
// 	fmt.Println("\"")
// 	fmt.Println(`"`)
// }
// func main() {
// 	t,f := true,false
// 	fmt.Printf("%T %v %t\n",t,t,t)
// 	fmt.Printf("%T %v %t\n",f,f,f)

// 	fmt.Println(true && true)
// 	fmt.Println(true && false)
// 	fmt.Println(false && false)
// 	fmt.Println(true || true)
// 	fmt.Println(true || false)
// 	fmt.Println(false || false)

// 	fmt.Println(!true)
// 	fmt.Println(!false)
// }

// func main(){
// 	var x int = 1
// 	xx:=float64(x)
// 	fmt.Printf("%T %v %f\n",xx,xx,xx)

// 	var y float64 = 1.2
// 	yy:=int(y)
// 	fmt.Printf("%T %v %d\n",yy,yy,yy)

// 	var s string = "14"
// 	i, _ := strconv.Atoi(s)
// 	fmt.Printf("%T %v\n",i,i)

// 	h := "Hello, World!"
// 	fmt.Println(string(h[0]))
// }

// func main() {
// 	var a [2]int
// 	a[0] = 100
// 	a[1] = 200
// 	fmt.Println(a)
// 	/*
// 	var b [2]int = [2]int{100, 200}
// 	b = append(b, 300) // エラー
// 	fmt.Println(b)
// 	*/

//		var b []int = []int{100, 200}
//		b = append(b, 300)
//		fmt.Println(b)
//	}
// func main() {
// 	n := []int{1, 2, 3, 4, 5, 6}
// 	fmt.Println(n)
// 	fmt.Println(n[2])
// 	fmt.Println(n[2:4])
// 	fmt.Println(n[:2])
// 	fmt.Println(n[2:])
// 	fmt.Println(n[:])

// 	n[2] = 100
// 	fmt.Println(n)

// 	var board = [][]int{
// 		[]int{0, 1, 2},
// 		[]int{3, 4, 5},
// 		[]int{6, 7, 8},
// 	}
// 	fmt.Println(board)

// 	n = append(n, 100, 200, 300, 400)
// 	fmt.Println(n)
// }

// func main() {
// 	n := make([]int, 3, 5)
// 	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
// 	n = append(n, 0, 0)
// 	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)
// 	n = append(n, 1, 2, 3, 4, 5)
// 	fmt.Printf("len=%d cap=%d value=%v\n", len(n), cap(n), n)

// 	a := make([]int, 3)
// 	fmt.Printf("len=%d cap=%d value=%v\n", len(a), cap(a), a)

// 	b := make([]int, 0)
// 	var c []int
// 	fmt.Printf("len=%d cap=%d value=%v\n", len(b), cap(b), b)
// 	fmt.Printf("len=%d cap=%d value=%v\n", len(c), cap(c), c)

// 	c = make([]int,5)
// 	// c = make([]int,0,5)
// 	for i:=0 ; i < 5; i++{
// 		c = append(c, i)
// 		fmt.Println(c)
// 	}
// 	fmt.Println(c)
// }

// func main() {
// 	m := map[string]int{"apple": 100, "banana": 200}
// 	fmt.Println(m)
// 	fmt.Println(m["apple"])
// 	m["banana"] = 300
// 	fmt.Println(m)
// 	m["new"] = 500
// 	fmt.Println(m)
// 	fmt.Println(m["nothing"])
// 	v, ok := m["apple"]
// 	fmt.Println(v, ok)
// 	v2, ok2 := m["nothing"]
// 	fmt.Println(v2, ok2)

// 	m2 := make(map[string]int)
// 	m2["pc"] = 5000
// 	fmt.Println(m2)

// 	/*
// 	var m3 map[string]int
// 	m3["pc"] = 5000
// 	fmt.Println(m3)
// 	*/

// 	// varで定義した場合はnil
// 	var s []int
// 	if s == nil{
// 		fmt.Println("Nil")
// 	}
// }

// func main() {
// 	b := []byte{72, 73}
// 	fmt.Println(b)
// 	fmt.Println(string(b))

// 	c := []byte("HI")
// 	fmt.Println(c)
// 	fmt.Println(string(c))
// }

// func add(x, y int) (int,int) {
// 	return x + y,x-y
// }

// func cal(price, item int) (result int) {
// 	result = price * item
// 	return
// }

// func main() {
// 	r1,r2 := add(10, 20)
// 	fmt.Println(r1,r2)

// 	r3 := cal(100, 2)
// 	fmt.Println(r3)

// 	f := func(x int){
// 		fmt.Println("inner func",x)
// 	}
// 	f(1)
// 	func(x int){
// 		fmt.Println("inner func",x)
// 	}(1)
// }

// func incrementGenerator() func() int {
// 	x:= 0
// 	return func() int {
// 		x++
// 		return x
// 	}
// }

// func circleArea(pi float64) func(radius float64) float64 {
// 	return func(radius float64) float64 {
// 		return pi * radius * radius
// 	}
// }

// func main() {
// 	counter := incrementGenerator()
// 	fmt.Println(counter())
// 	fmt.Println(counter())
// 	fmt.Println(counter())
// 	fmt.Println(counter())

// 	c1 := circleArea(3.14)
// 	fmt.Println(c1(2))
// 	c2 := circleArea(3)
// 	fmt.Println(c2(2))
// }

// func foo(params ...int){
// 	fmt.Println(len(params),params)
// 	for _, param := range params{
// 		fmt.Println(param)
// 	}
// }
// func main(){
// 	foo()
// 	foo(10,20)
// 	foo(10,20,30)
// 	s := []int{1,2,3}
// 	foo(s...)
// }


func main() {
	f := 1.11
	println(int(f))
	s := []int{1,2,5,6,2,3,1}
	fmt.Println(s[2:4])
	m := map[string]int{"Mike": 20, "Nancy": 24, "Messi": 30}
	fmt.Println(m)
}