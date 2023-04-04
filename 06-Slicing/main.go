package main

import "fmt"

func main() {
	x := [4]string{"red","white","blue","gray"}
	y := x[1:3]  // red white
	
	a := x[:3]  // red, white
	b := x[2:]  // blue, gray
	c := x[1:3] // white blue
	d := y[1:]  // blue  (slice of a slice)
	
	fmt.Println("x ", x)
	fmt.Println("y ", y)
	fmt.Println("a ", a)
	fmt.Println("b ", b)
	fmt.Println("c ", c)
	fmt.Println("d ", d)
	
}