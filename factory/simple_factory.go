package main

import "fmt"

type Book interface {
	Name() string
}

type chineseBook struct {
	name string
}

func (c *chineseBook) Name() string {
	return c.name
}

type mathBook struct {
	name string
}

func (c *mathBook) Name() string {
	return c.name
}

type englishBook struct {
	name string
}

func (c *englishBook) Name() string {
	return c.name
}

type physicalBook struct {
	name string
}

func (p *physicalBook) Name() string {
	return p.name
}

type Paper interface {
	Name()
}

type Person interface {
	GetBook(string) Book
}

type chineseAssistant struct {

}

func (c *chineseAssistant) GetBook(name string) Book {
	if name == "语文" {
		return &chineseBook{
			name: name,
		}
	}
	return nil
}

type mathAssistant struct {

}

func (m *mathAssistant) GetBook(name string) Book {
	if name == "数学" {
		return &mathBook{
			name : name,
		}
	}
	return nil
}

type englishAssistant struct {

}

func (e *englishAssistant) GetBook(name string) []Book {
	if name == "英语" {
		var englishBooks []Book
		englishBooks = append(englishBooks, &englishBook{
			name : name,
		})
		englishBooks = append(englishBooks, &englishBook{
			name : name,
		})
		englishBooks = append(englishBooks, &englishBook{
			name : name,
		})
		englishBooks = append(englishBooks, &englishBook{
			name : name,
		})
		return englishBooks
	}

	return nil
}

type physicalAssistant struct {

}

func (e *physicalAssistant) GetBook(name string) Book {
	if name == "物理" {
		return &physicalBook{
			name : name,
		}
	}
	return nil
}


func main() {
	c := &chineseAssistant{}
	fmt.Printf("小明领取了一本 %s 书\n", c.GetBook("语文").Name())
	m := &mathAssistant{}
	fmt.Printf("小明领取了一本 %s 书\n", m.GetBook("数学").Name())
	e := &englishAssistant{}
	fmt.Printf("小明领取了一本 %s 书\n", e.GetBook("英语").Name())
	p := &physicalAssistant{}
	fmt.Printf("小明领取了一本 %s 书\n", p.GetBook("物理").Name())
}