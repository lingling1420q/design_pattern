package main

import (
	"./monitor"
	"fmt"
	"log"
	"sync"
)

type Member interface {
	AskForHelp(string)
}

type member struct {
	name string
}

func (mem *member) AskForHelp(name string) {
	mo := monitor.GetMonitor(name)
	fmt.Printf("%s 寻求了 %s 的帮助 \n", mem.name, name)
	mo.Help(mem.name)
}

//func main() {
//	zhangsan := &member{
//		name : "张三",
//	}
//	zhangsan.AskForHelp("小明")
//
//	lisi := &member{
//		name : "李四",
//	}
//	lisi.AskForHelp("小明")
//
//	wangwu := &member {
//		name : "王五",
//	}
//	wangwu.AskForHelp("小明")
//
//
//
//}


func main() {

	var wg sync.WaitGroup

	wg.Add(3)

	go func() {
		zhangsan := &member {
			name : "张三",
		}
		zhangsan.AskForHelp("小红")
		defer wg.Done()
	}()

	go func() {
		lisi := &member {
			name : "李四",
		}
		lisi.AskForHelp("小明")
		defer wg.Done()
	}()

	go func() {
		wangwu := &member{
			name : "王五",
		}
		wangwu.AskForHelp("小王")
		defer wg.Done()
	}()

	wg.Wait()

	log.Println("sss")

}








