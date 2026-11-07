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
}
