package main

//type Template interface {
//	GoToSchool()
//	MorningReading()
//	TakeClasses()
//	ClassIsOver()
//	SchoolIsOver()
//}
//
//type template struct {
//	name string
//}
//
//func (t *template) GoToSchool() {
//	fmt.Println(t.name, "去学校")
//}
//
//func (t *template) MorningReading() {
//	fmt.Println(t.name, "晨读")
//}
//
//func (t *template) TakeClasses() {
//	fmt.Println(t.name, "上课")
//}
//
//func (t *template) ClassIsOver() {
//	fmt.Println(t.name, "下课")
//}
//
//func (t *template) SchoolIsOver() {
//	fmt.Println(t.name, "放学")
//}
//
//func (t *template) SetName (name string) {
//	t.name = name
//}
//
//func SchoolDay(t Template) {
//	t.GoToSchool()
//	t.MorningReading()
//	t.TakeClasses()
//	t.ClassIsOver()
//	t.SchoolIsOver()
//}
//
//type Student interface {
//	SchoolDay()
//}
//
//type student struct {
//	template
//}
//
//func (s *student) TakeClasses() {
//	fmt.Println(s.name,"上语文课")
//	fmt.Println(s.name,"上数学课")
//	fmt.Println(s.name,"上英语课")
//}
//
//func main() {
//	s := &student{}
//	s.SetName("小明")
//	SchoolDay(s)
//}