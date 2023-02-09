package main

import (
	"fmt"
)

type MyInterface interface{}

func main() {
	var n *int
	test(n)
}

func test(r interface{}) {
	switch value := r.(type) {
	case *int:
		if value == nil {
			fmt.Println("Success!")
		}

	}
}
