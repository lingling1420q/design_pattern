package main

//import "fmt"
//
//type Student interface {
//	selectThePrize(Prize)
//}
//
//type Prize interface {
//	Name() string
//}
//
//
//type student struct {
//	name string
//}
//
//type pen struct {
//	name string
//}
//
//func (p *pen) Name() string {
//	return "钢笔"
//}
//
//type basketball struct {
//	name string
//}
//
//func (b *basketball) Name() string {
//	return "篮球"
//}
//
//type notebook struct {
//	name string
//}
//
//func (n *notebook) Name() string {
//	return "本子"
//}
//
//
//
//func (s *student) selectThePrize(prize Prize) {
//	if prize == nil {
//		fmt.Println(s.name, "什么都没有选择")
//		return
//	}
//	fmt.Println(s.name, "选择了奖品 ：", prize.Name())
//}
//
//
//
//
//func main() {
//	xiaowang := &student {
//		name : "小王",
//	}
//	xiaowang.selectThePrize(&basketball{})
//
//	xiaohong := &student {
//		name : "小红",
//	}
//	xiaohong.selectThePrize(&pen{})
//
//	xiaoming := &student {
//		name : "小明",
//	}
//	xiaoming.selectThePrize(&notebook{})
//
//	xiaozhang := &student {
//		name : "小张",
//	}
//	xiaozhang.selectThePrize(nil)
//}

