package main

import (
	"fmt"
	"reflect"
)

func main() {
	x := [4]string{"red","white","blue","gray"}
	y := x[1:3]  // By Reference. y is an slice.
	fmt.Println(y, y[0], y[1])
	fmt.Println(reflect.TypeOf(x))
	fmt.Println(reflect.TypeOf(y))
	y[0] = "pink"
	fmt.Println(x)
	fmt.Println(y)
}