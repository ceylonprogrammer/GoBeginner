package main

import "fmt"

func main() {
	//standard for loop
	var i int = 0
	for i <= 5 {
		fmt.Println(i)
		i = i + 1
	}
	//shorthand version
	//initial ;condition ;increment
	fmt.Println("value of k is")
	for k := 1; k <= 8; k++ {
		fmt.Println(k)
	}
	//without condition loop will continue if not break statement
	for {
		fmt.Println("hello") //this goes infinite
		break
	}
	//Using continue keyword you can start net itaration
	fmt.Println("continue keyword")
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)

	}
	//loop through string
	const name string = "sameera"
	for i := 0; i < len(name); i++ {
		fmt.Printf("%c ", name[i])
	}
	fmt.Println("lenghth os a name ", len(name))

}
