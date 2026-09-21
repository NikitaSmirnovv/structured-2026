package main

import "fmt"

var age uint = 20

func printnameandage(name string, age uint) {
	fmt.Println("hello world")
	fmt.Println("you are", age, "years old")
}

func main() {

	var name string
	var age uint

	fmt.Println("Enter your name")
	fmt.Scan(&name)

	fmt.Println("Enter your age")
	fmt.Scan(&age)

	printnameandage(name, 40)
}
