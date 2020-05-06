package main

import (
	"fmt"
)

type teacher interface {
	arrangeHomework()
}

type mathTeacher struct {
	students []student
}

func (m *mathTeacher) arrangeHomework() {
	fmt.Println("数学老师布置作业")

	for _, s := range m.students {
		s.writingMathHomework()
	}
}

type physicsTeacher struct {
	students []student
}

func (p *physicsTeacher) arrangeHomework() {
	fmt.Println("物理老师布置作业")

	for _, s := range p.students {
		s.writingPhysicsHomework()
	}
}

type student interface {
	writingPhysicsHomework()
	writingMathHomework()
}

type xiaoming struct {

}

func (ming *xiaoming) writingMathHomework() {
	fmt.Println("小明做 100 道数学题")
}

func (ming *xiaoming) writingPhysicsHomework() {
	fmt.Println("小明做 50 道物理题")
}

type xiaohong struct {

}

func (hong *xiaohong) writingMathHomework() {
	fmt.Println("小红做 50 道数学题")
}

func (hong *xiaohong) writingPhysicsHomework() {
	fmt.Println("小红做 50 道物理题")
}


func main() {
	var students []student
	students = append(students, &xiaoming{})
	students = append(students, &xiaohong{})

	m := &mathTeacher{
		students : students,
	}
	p := &physicsTeacher{
		students : students,
	}
	m.arrangeHomework()
	p.arrangeHomework()
}