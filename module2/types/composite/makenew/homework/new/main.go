package main

import "fmt"

func main() {
	var n *int = new(int)
	fmt.Println(*n)
}
