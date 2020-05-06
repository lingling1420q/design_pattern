package main

import "fmt"

type Person interface  {
	SetName(string)
	GetName() string
}

type person struct {
	name string
}

func (p *person) GetName() string {
	return p.name
}

func (p *person) SetName(name string) {
	p.name = name
}

func main() {
	p := &person{
		name : "张三",
	}

	fmt.Println(p.GetName())

	doSomething(p)
}

func doSomething(p Person) {
	fmt.Println(p.GetName())
}
