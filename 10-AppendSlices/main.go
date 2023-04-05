package main

import "fmt"

func main() {
	colors := []string{"red","white","green"}
	lightColors := []string{"orange","yellow"}
	darkColors := []string{"black","brown","blue"}
	
	fmt.Println("colors",colors)
	
	colors = append(colors,lightColors...)
	colors = append(colors,darkColors[1:3]...)
	
	fmt.Println("colors",colors)
	buildLightColors()
}

func buildLightColors(){
	colors := []string{"red","white","green"}
	lightColors := make([]string,0)
	lightColors = append(lightColors,colors[:2]...)
	fmt.Println("lightColors",lightColors)
}