package main

import "fmt"

func main() {

	//explicit variable declaration, specifically said what type it will be, better practice to play it safe
	var number uint16 = 260
	//implicit variable
	var number2 = 255
	number = number + 5
	//walrus symbol is the same as writing var number3 uint16 = 6
	number3 := 6

	var thisFloat float32
	var thisString string
	fmt.Println("Hello World!", number, number2, number3)
	fmt.Println(fmt.Printf("%T", thisFloat))
	fmt.Printf("%T", thisString)

}
