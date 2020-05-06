package main

import (
	"fmt"
	"net/http"
)

type Student interface {
	SchoolDay()
}

type xiaoming struct {
	template
	name string
}

type xiaozhang struct {
	template
	name string
}

func (xiaozhang *xiaozhang) TakingClasses() {
	fmt.Println(xiaozhang.name, "上课")
	fmt.Println(xiaozhang.name, "补习")
}

type Template interface {
	GoToSchool ()
	MorningReading()
	TakingClasses()
	ClassIsOver()
	SchoolIsOver()
}

type template struct {
	name string
}

func (s *template) GoToSchool () {
	fmt.Println(s.name, "去学校")
}

func (s *template) MorningReading () {
	fmt.Println(s.name, "晨读")
}

func (s *template) TakingClasses () {
	fmt.Println(s.name, "去上课")
}

func (s *template) ClassIsOver () {
	fmt.Println(s.name, "下课")
}

func (s *template) SchoolIsOver () {
	fmt.Println(s.name, "放学")
}

func (s *template) SetName(name string) {
	s.name = name
}

func SchoolDay(s Template) {
	s.GoToSchool()
	s.MorningReading()
	s.TakingClasses()
	s.ClassIsOver()
	s.SchoolIsOver()
}

func main() {
	xiaoming := &xiaoming{
		name : "小明",
	}
	xiaoming.SetName(xiaoming.name)
	SchoolDay(xiaoming)

	xiaozhang := &xiaozhang{
		name : "小张",
	}
	xiaozhang.SetName(xiaozhang.name)
	SchoolDay(xiaozhang)

	http.Handle()
}


