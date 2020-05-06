package main
//
//import (
//	"fmt"
//	"log"
//)
//
//type Monitor interface {
//	Help(string)
//}
//
//type Member interface {
//	AskForHelp()
//}
//
//type monitor struct {
//	name string
//}
//
//func (m *monitor) Help(name string) {
//	fmt.Printf("%s 帮助了 %s\n", m.name, name)
//}
//
//type member struct {
//	name string
//}
//
//func (z *member) AskForHelp() {
//	m := &monitor{
//		name : "小明",
//	}
//	m.Help(z.name)
//}
//
//
//func main() {
//	z := &member{
//		name : "张三",
//	}
//	l := &member{
//		name : "李四",
//	}
//
//	z.AskForHelp()
//	l.AskForHelp()
//
//	log.Println("sss")
//}