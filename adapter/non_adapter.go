package main

import "fmt"

type Backpack interface {
	putTheBook([]string)
}

type backpack struct {

}

func (bp *backpack) putTheBook(books []string) {
	for _, book := range books {
		fmt.Printf("把 %s 放入书包里面 \n", book)
	}
}

type TargetBackpack interface {
	putTheBook([]string, []string)
}

type Adapter interface {
	putTheBook([]string, []string)
}

type adapter struct {
	bp backpack
}

func (ad *adapter) putTheBook(books []string, papers []string) {
	ad.bp.putTheBook(books)

	for _, paper := range papers {
		fmt.Printf("把 %s 放入书包夹层里面 \n", paper)
	}
}


func main() {
	books := []string{"语文书","数学书","英语书"}
	papers := []string{"语文试卷","数学试卷"}

	bp := &backpack{}
	ad := &adapter{}

	bp.putTheBook(books)

	fmt.Println("===========================")
	ad.putTheBook(books, papers)


}

