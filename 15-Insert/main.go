package main

import "fmt"

func main() {
	myColors := []string{"red","white","blue","green","orange"}
	fmt.Println(myColors)
	myColors = append(myColors[0:2],append([]string{"black","pink"},myColors[2:]...)...)
	fmt.Println(myColors)
	
	nums := []int{10,40,50,60}
	fmt.Println(nums)
	moreNums := []int{20,30}
	nums = append(nums[:1],append(moreNums,nums[1:len(nums)-1]...)...)
	fmt.Println(nums)
	
	firstNumbers := []int{1,2,3,4}
	secondNumbers := []int{5,6,7,8}
	thirdNumbers := []int{9,10,11}
	allNumbers := append(append(firstNumbers,secondNumbers...),thirdNumbers...)
	fmt.Println(allNumbers)
		
}