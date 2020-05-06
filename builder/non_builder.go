package main
//
//import "fmt"
//
//type class struct {
//	monitor string
//	studentAssistant string
//	chineseAssistant string
//	mathAssistant string
//	englishAssistant string
//}
//
//func main() {
//	c := &class {}
//	c.setMonitor("小明")
//	c.setStudentAssistant("小红")
//
//
//
//	fmt.Printf("班长 ： %s, 学委 ： %s", c.monitor, c.studentAssistant)
//}
//
//
//func (c *class) setMonitor(monitor string) {
//	fmt.Println("自主报名...")
//	fmt.Println("班级同学投票...")
//	fmt.Println("统计票数，得出结果...")
//	c.monitor = monitor
//}
//
//
//
//func (c *class) setStudentAssistant(studentAssistant string) {
//	fmt.Println("自主报名...")
//	fmt.Println("老师直接指定...")
//	c.studentAssistant = studentAssistant
//}