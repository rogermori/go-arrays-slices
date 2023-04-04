package main

import "fmt"

func main() {
	colors := []string{"red","yellow","green"}
	fmt.Println("colors",colors, len(colors), cap(colors))
	colors = append(colors,"black","brown")
	fmt.Println("colors",colors, len(colors), cap(colors))
}