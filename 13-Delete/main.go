package main

import "fmt"

func main() {
	myColors := []string{"red","white","blue","green","orange"}
	fmt.Println(myColors,len(myColors),cap(myColors))
	myColors = append(myColors[0:2],myColors[3:]...)
	fmt.Println(myColors,len(myColors),cap(myColors))
}