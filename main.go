package main

import (
	"fmt"
)

func main() {
	x := [4]string{"red","white","blue","gray"}
	fmt.Println("x = ", x)
	y := x
	fmt.Println("y = ", y)
	x[0] = "pink"
	fmt.Println("x = ", x)
	fmt.Println("y = ", y)
}