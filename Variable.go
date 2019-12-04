package main

import "fmt"

func main() {
	var name = "sameera"
	//var declare one or more variable
	fmt.Println(name)
	var num1, num2 int = 1, 4 //can declare multiple variable with single line
	fmt.Println(num1, num2)
	var status = true
	fmt.Println(status) //go aknowledge with the data type
	var value = 7.0
	fmt.Println(value) //don't need to explicit about the data type

	//variable can initializing without declaring the value and initial value is zero value
	var nullValueVariable int
	fmt.Println(nullValueVariable)

	//shorthand typing of a variable
	degree := "IT"
	fmt.Println(degree)
}
