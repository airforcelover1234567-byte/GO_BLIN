package main

import "fmt"

// for->Only loop availabe in golang

func main() {
	//while loop-As there is no while loops in go, we can use a for loop as a while loop
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i++
	}
	// Infinite loop
	// for {
	// 	fmt.Println("Rounak is so handsome")
	// }

	//For loop
	for i := 0; i <= 5; i++ {
		fmt.Println(i)
		if i == 2 {
			fmt.Println("Good luck in heaven")
			break
		}
		if i == 1 {
			continue
		}
	}
}
