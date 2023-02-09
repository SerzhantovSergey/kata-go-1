package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	s = Append(s)
	fmt.Print(s)
}

func Append(s []int) []int {

	s = append(s, 5)
	return s
}
