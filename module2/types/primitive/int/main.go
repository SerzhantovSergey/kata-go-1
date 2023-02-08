package main

import "fmt"

func main() {
	var uintNumber uint16 = 1 << 15
	var from = int16(uintNumber)
	uintNumber--
	var to = int16(uintNumber)
	fmt.Println(from, to)
}
