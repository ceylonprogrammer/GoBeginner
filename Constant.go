package main

//constant declare the constant value
import "fmt"

var myage int

const age int = 31

func main() {
	const name string = "Sameera"
	fmt.Println(name)
	fmt.Println(age)
	//can not assign age again because it is constant
	myage = 45
	myage = 55
	fmt.Println(myage)

}
