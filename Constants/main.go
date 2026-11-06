package main

import "fmt"

const age = 24

func main() {
	const name string = "Rounak"
	fmt.Println(age)
	fmt.Println(name)

	const (
		name1 = "Rounak"
		name2 = "Sounak"
	)

	fmt.Println(name1, name2)
}
