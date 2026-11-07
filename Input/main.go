package main

import "fmt"

func main() {
	var name string
	var age int32

	fmt.Println("Enter your name")
	fmt.Scanln(&name)

	fmt.Println("Enter your age")
	fmt.Scanln(&age)

	fmt.Printf("Your name is %s .You are %d years old\n", name, age)
}
