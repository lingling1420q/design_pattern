package main

import "fmt"

type Student interface {
	selectThePrize(Prize)
}

type Prize interface {
	Name() string
}

type student struct {
	name string
}

func (s *student) selectThePrize(p Prize) {
	fmt.Println(s.name, "选择了奖品 : ", p.Name())
}

type pen struct {

}

type watch struct {

}

func (w *watch) Name() string {
	return "手表"
}

func (p *pen) Name() string {
	return "钢笔"
}

type basketball struct {

}

func (b *basketball) Name() string {
	return "篮球"
}

type notebook struct {

}

func (n *notebook) Name() string {
	return "本子"
}

func main() {
	xiaowang := &student{
		name : "小王",
	}
	xiaowang.selectThePrize(&watch{})

	xiaohong := &student{
		name : "小红",
	}
	xiaohong.selectThePrize(&notebook{})

	xiaoming := &student{
		name : "小明",
	}
	xiaoming.selectThePrize(&basketball{})
}










