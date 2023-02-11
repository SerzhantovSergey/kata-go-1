package main

import "fmt"

func main() {
	var n *int
	n = new(int)
	fmt.Println(*n)
}
