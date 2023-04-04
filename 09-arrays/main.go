package main

import "fmt"

func main() {
	alphabet := "abcdefghijk"
	alpha := alphabet[:4]
	a     := alphabet[0]
	fmt.Println(alphabet)
	fmt.Println(alpha)
	fmt.Println(a, string(a))
}