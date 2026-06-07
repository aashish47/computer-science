package main

import "fmt"

type User struct {
	name string
	age  int
}

func (u *User) setName(name string) {
	u.name = name
}

func (u *User) setAge(age int) {
	u.age = age
}

func (u *User) String() string {
	return fmt.Sprintf("name: %v, age: %v", u.name, u.age)
}

func main() {
	aashish := &User{
		name: "aashish",
		age:  29,
	}
	fmt.Println(aashish)
	aashish.setAge(22)
	fmt.Println(aashish)
}
