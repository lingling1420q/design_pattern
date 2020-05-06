package main

//import (
//	"fmt"
//)
//
//type teacher struct {
//	students []student
//}
//
//func (t *teacher) arrangeHomework() {
//	fmt.Println("数学老师布置 100 道数学题")
//
//	for _, s := range t.students {
//		s.writingHomework()
//	}
//}
//
//type student interface {
//	writingHomework()
//}
//
//type xiaoming struct {
//
//}
//
//func (ming *xiaoming) writingHomework() {
//	fmt.Println("我是小明，我需要完成 100 道数学题")
//}
//
//type xiaohong struct {
//
//}
//
//func (hong *xiaohong) writingHomework() {
//	fmt.Println("我是小红，我需要完成 50 道数学题")
//}
//
//
//
//
//func main() {
//	var students []student
//	students = append(students, &xiaoming{})
//	students = append(students, &xiaohong{})
//
//	t := &teacher{students: students}
//	t.arrangeHomework()
//}