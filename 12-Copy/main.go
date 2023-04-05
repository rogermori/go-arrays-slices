package main

import "fmt"

func main() {
	// copy(target, source)  passed data by value
	names := []string{"Ana","Rose","Tony","Mary"}
	relatives := make([]string,2)
	copy(relatives,names)
	
	fmt.Println("names",names)
	fmt.Println("relatives",relatives)
	
	copyIntegers()
}

func copyIntegers(){
	nums := []int{33,45,56}
	n := make([]int,8)
	copy(n,nums)
	
	fmt.Println("nums",nums)
	fmt.Println("n",n)
}