package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func newPerson() *Person {
	p:= person{
		Name: "Liam",
		Age:  30,
	}
	return &p
}

func main() {
	p := newPerson()
	fmt.Printf("Name: %s, Age: %d\n", p.Name, p.Age)
	}

	dog := struct {
		Breed string
		Age   int
	}{
		Breed: "Labrador",
		Age:   5,
	}
	fmt.Printf("Breed: %s, Age: %d\n", dog.Breed, dog.Age)
}