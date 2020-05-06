package main

//import "fmt"
//
//type Book interface {
//	Name() string
//}
//
//type ChineseBook struct {}
//
//func (c *ChineseBook) Name() string {
//	return "语文"
//}
//
//type MathBook struct {}
//
//func (m *MathBook) Name() string {
//	return "数学"
//}
//
//type EnglishBook struct {}
//
//func (e *EnglishBook) Name() string {
//	return "英语"
//}
//
//
//
func main() {
	c := &ChineseBook{}
	fmt.Printf("小明领了%s课本\n", c.Name())
	m := &MathBook{}
	fmt.Printf("小明领了%s课本\n", m.Name())
	e := &EnglishBook{}
	fmt.Printf("小明领了%s课本\n", e.Name())
}