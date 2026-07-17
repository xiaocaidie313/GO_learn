package main

import "fmt"

type Hero struct {
	Name  string
	Ad    string
	level string
}

func (h Hero) GetName() {
	fmt.Println("Name =", h.Name)
}

type Human struct {
	Hero
	leg int
}

func main() {
	hero := Hero{Name: "张飞", Ad: "燕人"}
	hero.GetName()

	human := Human{Hero: hero, leg: 2}
	fmt.Printf("Human: %+v\n", human)

	// interface 多态（见 interface.go）
	Talk(Dog{Name: "旺财"})
	Talk(Cat{Name: "咪咪"})
	myFunc("hello")
}
