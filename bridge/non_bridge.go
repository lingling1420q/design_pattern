package main

import "fmt"

type Class interface {
	Number() int
}

type Clothes interface {
	Color() string
}

type classOne struct {
	clo Clothes
}

func (c *classOne) Number() int {
	return 1
}

type classTwo struct {

}

func (c *classTwo) Number() int {
	return 2
}

type classThree struct {}

func (c *classThree) Number() int {
	return 3
}

type yellowClothes struct {}

func (y *yellowClothes) Color() string {
	return "黄色"
}

type whiteClothes struct {}

func (y *whiteClothes) Color() string {
	return "白色"
}

type redClothes struct{}

func (y *redClothes) Color() string {
	return "红色"
}

func wearClothes() {
	class := &classTwo{}
	clothes := &redClothes{}
	fmt.Println("小明是", class.Number(),"班的同学")
	fmt.Println("小明穿", clothes.Color(), "的衣服")
}

func main() {
	wearClothes()
}