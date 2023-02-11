package main

import "fmt"

type User struct {
	ID   int
	Name string
}

func (u User) GetName(username string) string {
	return u.Name
}

type Userer interface {
	GetName(string) string
}

func main() {
	var i Userer = User{}
	_ = i
	fmt.Println("Success!")
}
