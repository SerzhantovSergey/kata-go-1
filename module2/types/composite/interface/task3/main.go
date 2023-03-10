package main

import (
	"fmt"
)

type User struct {
	ID   int
	Name string
}

func (u *User) GetName() string {
	return u.Name
}

type Userer interface {
	GetName() string
}

func main() {

	u := []User{
		{
			ID:   34,
			Name: "Annet",
		},
		{
			ID:   55,
			Name: "John",
		},
		{
			ID:   89,
			Name: "Alex",
		},
	}
	v := make([]Userer, len(u))
	for i := range u {
		v[i] = &u[i]
	}

	users := v
	testUserName(users)
}

func testUserName(users []Userer) {
	for _, u := range users {
		fmt.Println(u.GetName())
	}
}
