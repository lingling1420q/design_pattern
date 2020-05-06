package main

import "fmt"

type Student interface {
	HandleInHomework(*teacher)
}

type Teacher interface {
	CollectHomework(string)
}

type student struct {
	name string
}

func (s *student) HandleInHomework(a *assistent) {
	fmt.Println(s.name , "交作业给", a.name)
	a.CollectHomework(s.name)
}

type teacher struct {
	name string
}

func (t *teacher) CollectHomework(name string) {
	fmt.Println(t.name, "收到了", name, "的作业")
}

type assistent struct {
	name string
	t *teacher
}

func (a *assistent) CollectHomework(name string) {
	fmt.Println(a.name, "收到了", name, "的作业")
	fmt.Println(a.name,"统计谁交了作业")
	a.t.CollectHomework(name)
}

func main() {
	zhangsan := &student{
		name :"张三",
	}
	t := &teacher{
		name : "语文老师",
	}
	a := &assistent{
		name : "语文课代表",
		t : t,
	}
	zhangsan.HandleInHomework(a)
}


