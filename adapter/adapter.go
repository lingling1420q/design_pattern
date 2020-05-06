package main

//import "fmt"
//
//type Backpack interface {
//	putTheBook([]string)
//}
//
//type backpack struct {}
//
//func (b *backpack) putTheBook(books []string) {
//	for _, book := range books {
//		fmt.Printf("把 %s 放到书包里 \n", book)
//	}
//}
//
//type TargetBack interface {
//	putTheBook([]string, []string)
//}
//
//type targetBack struct{}
//
//func (t *targetBack) putTheBook(books []string, papers []string) {
//	for _, book := range books {
//		fmt.Printf("把 %s 放到书包里 \n", book)
//	}
//	for _, paper := range papers {
//		fmt.Printf("把 %s 放到书包夹层里 \n", paper)
//	}
//}
//
//type Adapter interface {
//	putTheBook([]string, []string)
//}
//
//type adapter struct {
//	bp Backpack
//}
//
//func (a *adapter) putTheBook(books []string, papers []string) {
//	a.bp.putTheBook(books)
//
//	for _, paper := range papers {
//		fmt.Printf("把 %s 放到书包夹层里 \n", paper)
//	}
//
//}
//
//func main() {
//	books := []string{"语文书","数学书","英语书"}
//	papers := []string{"语文试卷","数学试卷","英语试卷"}
//	bp := &backpack{}
//	ad := &adapter{
//		bp : bp,
//	}
//
//	bp.putTheBook(books)
//	fmt.Println("===================================")
//	ad.putTheBook(books, papers)
//}
//



