package main

import "fmt"

func main() {

	//conditions and boolean expressions
	x := 5
	// y := 6
	val := x == 5
	fmt.Printf("%t \n", val)

	//basic arthmetic
	// var num1 int = 8
	// var num2 int = 4
	// answer := num1 / num2
	// fmt.Printf("%d \n", answer)
	//can pull in "math"

	// // console input and type conversion
	// scanner := bufio.NewScanner(os.Stdin)
	// fmt.Printf("Type the year you were born: ")
	// scanner.Scan()
	// //the underscore says if this doesnt work then just ignore it. ignore the error
	// input, _ := strconv.ParseInt(scanner.Text(), 10, 64)
	// fmt.Printf("You will be %d years old at the end of 2020", 2020-input)

	//explicit variable declaration, specifically said what type it will be, better practice to play it safe
	// var number uint16 = 260
	// //implicit variable
	// var number2 = 255
	// number = number + 5
	// //walrus symbol is the same as writing var number3 uint16 = 6
	// number3 := 6

	// var thisFloat float64
	// var thisString string
	// fmt.Println("Hello World!")
	// fmt.Println(fmt.Printf("%T", thisFloat))
	// fmt.Printf("%T", thisString)

	// fmt.Printf("%x", 100)

}
