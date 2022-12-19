package main

import "fmt"

func main() {
	var a = []int{0, 1, 1, 0, 0, 1}
	var b = []int{10, 2, 3, 1, 7, 18}

	fmt.Println(MinDifference(a, b))
}
