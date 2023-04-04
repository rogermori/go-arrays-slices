package main

import "fmt"

/*
- Declare an array (oddNumbers) of type int, with a length of 5.
- Assign numbers 1,2,3,4,5 to the array, at the time of declaration.
- Create a "for range" loop to traverse the array.
- Print only the odd numbers. As for each even number, print two dashes.
  Output:
	  1 -- 3 -- 5
*/

func main() {
	oddNumbers := [5]int {1, 2, 3, 4, 5}
	
	for index := range oddNumbers {
		if oddNumbers[index]%2 == 0 {
			fmt.Print("--")
		} else {
			fmt.Print(oddNumbers[index])
		}
	}
}