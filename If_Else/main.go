package main

import "fmt"

func main() {
	age := 17

	if age >= 18 {
		fmt.Println("Person is an adult")
	} else if age >= 12 {
		fmt.Println("You are a teenager")
	} else {
		fmt.Println("Person is a kid")
	}
}
