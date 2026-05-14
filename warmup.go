package main

import "fmt"

func main() {
	fmt.Println("Hello World!")

	var intNum int // int8 int 16 and so on
	fmt.Println(intNum)

	var floatNum float32 = 1.23
	fmt.Println(floatNum)

	var myString string = "Hello World"
	fmt.Println(myString)

	// boleans are similar
	// default number is 0, default boolean is false, string is empty

	myVar := "text" // drop the var keyword
	var1, var2 := 1, 2
	fmt.Println(myVar, var1, var2)

	// everything applies to constant, except you can't change its value

	var printValue string = "Hello World"
	printMe(printValue)

	var num int = 11
	var den int = 5
	var result, rem int = intDivision(num, den)
	fmt.Printf("The result is %v with remainder %v", result, rem)
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(num int, den int) (int, int) {
	var result int = num / den
	var rem int = num % den
	return result, rem
}

// go run
