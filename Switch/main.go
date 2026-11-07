package main

import (
	"fmt"
	"time"
)

func main() {
	//simple switch
	i := 2

	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("other")
	}

	//Multiple condition Switch
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Its a weekend")
	default:
		fmt.Println("It is a workday")
	}

	//type switch
	whoami := func(i interface{}) { //interface{} means any type
		switch i.(type) {
		case int:
			fmt.Println("integer")
		case float32:
			fmt.Println("float")
		case string:
			fmt.Println("string")
		default:
			fmt.Println("Other")
		}
	}
	whoami(2)
}
