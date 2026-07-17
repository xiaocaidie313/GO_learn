package main

import "fmt"

type Person struct {
	Name     string
	Age      int
	Address  string
	Salary   float64
	Birthday string
}

type PersonOption func(*Person)

func NewPerson(options ...PersonOption) Person {
	p := Person{}
	for _, op := range options {
		op(&p)
	}
	if p.Age <= 0 {
		p.Age = 0
	}
	return p
}

func WithName(name string) PersonOption {
	return func(p *Person) { p.Name = name }
}
func WithAge(age int) PersonOption {
	return func(p *Person) { p.Age = age }
}
func WithAddress(address string) PersonOption {
	return func(p *Person) { p.Address = address }
}

func main() {
	p := NewPerson(
		WithName("张三"),
		WithAge(25),
		WithAddress("北京"),
	)
	fmt.Printf("%+v\n", p)
}
