package main

import "fmt"

func main() {
	mySlice := []int{33,34,35}
	fmt.Println(mySlice,len(mySlice),cap(mySlice))
	mySlice = append(mySlice,make([]int,7)...)
	fmt.Println(mySlice,len(mySlice),cap(mySlice))
}