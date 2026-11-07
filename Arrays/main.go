package main

import "fmt"

// Numbered sequence of specific length
func main() {
	var nums [4]int
	nums[0] = 1
	fmt.Println(len(nums)) //len() helps us find the length of an array
	fmt.Println(nums)
	fmt.Println(nums[0])

	var vals [2]bool
	vals[1] = true
	fmt.Println(vals)

	var names [3]string
	names[0] = "Kutta"
	fmt.Println(names)
}
