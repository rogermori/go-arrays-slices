package main

import "fmt"

func main() {
	// make([]Type, length, capacity)
	names := make([]string, 5,20)
	primes := new([20]int)[0:5]
	colors := []string{"red","yellow","green"}
	
	fmt.Println("names", names, len(names), cap(names))
	fmt.Println("primes", primes, len(primes), cap(primes))
	fmt.Println("colors", colors, len(colors), cap(colors))
	
}