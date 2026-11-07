package main

import "fmt"

func main() {
	var n int
	fmt.Println("Enter the multiplication table you want till")
	fmt.Scanln(&n)

	for i := 1; i <= n; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Println(i * j)
		}
		fmt.Println("\n")
	}
}
