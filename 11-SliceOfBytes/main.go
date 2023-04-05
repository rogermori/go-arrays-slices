package main

import "fmt"

/*
Declare a slice named sliceOfBytes of type []byte, with a length of zero.

Append the following numbers as byte elements to sliceOfBytes: 
72,101,108,108,111

Using Println, display the contents of sliceOfBytes.

Using Println, along with the string() function,
display sliceOfBytes so that it shows the word Hello.
*/
func main() {
    sliceOfBytes := make([]byte,0)
	sliceOfBytes = append(sliceOfBytes,72,101,108,108,111)
	fmt.Println(sliceOfBytes)
	fmt.Print(string(sliceOfBytes))
}