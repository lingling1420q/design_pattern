package main


//import "fmt"
//
//type Book interface {
//	Name() string
//}
//
//type ChineseBook struct {
//	name string
//}
//
//func (c *ChineseBook) Name() string {
//	return c.name
//}
//
//
//type MathBook struct {
//	name string
//}
//
//func (c *MathBook) Name() string {
//	return c.name
//}
//
//type EnglishBook struct {
//	name string
//}
//
//func (c *EnglishBook) Name() string {
//	return c.name
//}
//
//type BookFactory interface {
//	GetBook(string) Book
//}
//
//type chineseAssistant struct {}
//
//func (c *chineseAssistant) GetBook(bookName string) Book {
//	return &ChineseBook{
//		name : bookName,
//	}
//
//}
//
//type mathAssistant struct {}
//
//func (c *mathAssistant) GetBook(bookName string) Book {
//	return &MathBook{
//		name: bookName,
//	}
//}
//
//type englishAssistant struct {}
//
//func (e *englishAssistant) GetBook(bookName string) Book {
//	return &EnglishBook{
//		name : bookName,
//	}
//}
//
//
//func main() {
//	s := &chineseAssistant {}
//	fmt.Printf("语文课代表领了一本%s书 \n",s.GetBook("语文").Name())
//
//	m := &mathAssistant {}
//	fmt.Printf("数学课代表领了一本%s书 \n",m.GetBook("数学").Name())
//
//	e := &englishAssistant {}
//	fmt.Printf("英语课代表领了一本%s书 \n",e.GetBook("英语").Name())
//}
