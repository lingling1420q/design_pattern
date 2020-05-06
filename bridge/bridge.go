package main

//import "fmt"
//
//type Clothes interface {
//	Color() string
//}
//
//type blackClothes struct {}
//
//func (b *blackClothes) Color() string {
//	return "黑"
//}
//
//type whiteClothes struct {}
//
//func (w *whiteClothes) Color() string {
//	return "白"
//}
//
//type redClothes struct {}
//
//func (r *redClothes) Color() string {
//	return "红"
//}
//
//type Class interface {
//	Number() int
//}
//
//type classOne struct {
//	clothes Clothes
//}
//
//func (c *classOne) Number() int {
//	return 1
//}
//
//type classTwo struct {
//	whiteClothes
//}
//
//func (c *classTwo) Number() int {
//	return 2
//}
//
//type classThree struct {
//	redClothes
//}
//
//func (c *classThree) Number() int {
//	return 3
//}
//
//func wearClothes() {
//	c := &classOne{}
//	c.clothes = &redClothes{}
//	fmt.Println(c.Number(),"班穿", c.clothes.Color(), "色的衣服")
//}
//
//func main() {
//	wearClothes()
//}
//
//
//
//
//
