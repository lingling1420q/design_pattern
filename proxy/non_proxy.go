package main

//import "fmt"
//
//type Student interface {
//	HandInHomework(*teacher)
//}
//
//type student struct {
//	name string
//}
//
//func (s *student) HandInHomework(t *teacher) {
//	fmt.Println(s.name, "交作业给 : ", t.name)
//	t.collectHomework(s.name)
//}
//
//
//type Teacher interface {
//	collectHomework(string)
//}
//
//
//type teacher struct {
//	name string
//}
//
//func (t *teacher) collectHomework(name string) {
//	fmt.Printf("%s 收到 %s 的作业 \n", t.name, name)
//}
//
//
//
//
//func main() {
//	zhangsan := &student{
//		name : "张三",
//	}
//	t := &teacher {
//		name : "语文老师",
//	}
//	zhangsan.HandInHomework(t)
//
//	lisi := &student {
//		name : "李四",
//	}
//	lisi.HandInHomework(t)
//}