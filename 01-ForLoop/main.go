package main

import "fmt"

func main() {
	colors := [4]string{"red","white","black","blue"}
	
	for i:= 0; i < len(colors); i++ {
		fmt.Println(colors[i])
	}
	
	for i := range colors {
		fmt.Println(colors[i])
	}
	
}