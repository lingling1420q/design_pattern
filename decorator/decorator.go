package main

import "fmt"

type Person interface {
	wearClothes()
}

type person struct {}

func (p *person) wearClothes() {
	fmt.Println("穿衣服")
	fmt.Println("穿裤子")
	fmt.Println("穿鞋子")
}


type student struct{
	p *person
}

func (s *student) wearClothes() {
	s.p.wearClothes()
	fmt.Println("背书包")
}

type primaryStudent struct {
	s *student
}

func (ps *primaryStudent) wearClothes() {
	ps.s.wearClothes()
	fmt.Println("戴红领巾")
}

func main() {
	p := &person{}
	fmt.Println(" person : ")
	p.wearClothes()

	s := &student{p : p}
	fmt.Println(" student : ")
	s.wearClothes()

	ps := &primaryStudent{ s : s}
	fmt.Println("primary student : ")
	ps.wearClothes()
}

