package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func newPerson() *Person {
	p := Person{Name: "Liam", Age: 30}
	return &p
}

func main() {
	p := newPerson()
	fmt.Printf("Person: %+v\n", p)

	// 匿名结构体
	dog := struct {
		Breed string
		Age   int
	}{Breed: "Labrador", Age: 5}
	fmt.Printf("Dog: %+v\n", dog)
}
