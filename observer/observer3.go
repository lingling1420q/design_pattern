package main

//import (
//	"fmt"
//)
//
//
//type teacher interface {
//	arrangeHomework()
//}
//
//type mathTeacher struct {
//	students []student
//}
//
//func (m *mathTeacher) arrangeHomework() {
//	fmt.Println("数学老师布置 100 道数学题")
//
//	for _, s := range m.students {
//		s.writingMathHomework()
//	}
//}
//
//type physicsTeacher struct {
//	students []student
//}
//
//func (p *physicsTeacher) arrangeHomework() {
//	fmt.Println("物理老师布置 50 道物理题")
//
//	for _, s := range p.students {
//		s.writingPhysicsHomework()
//	}
//}
//
//
//type student interface {
//	writingMathHomework()
//	writingPhysicsHomework()
//}
//
//type xiaoming struct {
//
//}
//
//func (ming *xiaoming) writingMathHomework() {
//	fmt.Println("我是小明，我需要完成 100 道数学题")
//}
//
//func (ming *xiaoming) writingPhysicsHomework() {
//	fmt.Println("我是小明，我需要完成 50 道物理题")
//}
//
//type xiaohong struct {
//
//}
//
//func (hong *xiaohong) writingMathHomework() {
//	fmt.Println("我是小红，我需要完成 50 道数学题")
//}
//
//func (hong *xiaohong) writingPhysicsHomework() {
//	fmt.Println("我是小红，我需要完成 50 道物理题")
//}
//
//
//
//func main() {
//	var students []student
//	students = append(students, &xiaoming{})
//	students = append(students, &xiaohong{})
//
//	m := &mathTeacher{students: students}
//	m.arrangeHomework()
//
//	p := &physicsTeacher{students: students}
//	p.arrangeHomework()
//}