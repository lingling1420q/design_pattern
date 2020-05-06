package main

//type student struct {
//	name string
//}
//
//
//func (s *student) Reading() int {
//	score := 40
//	fmt.Println("阅读理解得分为 : ", score)
//	return score
//}
//
//func (s *student) Listening() int {
//	score := 20
//	fmt.Println("听力得分为 : ", score)
//	return score
//}
//
//func (s *student) Writing() int {
//	score := 25
//	fmt.Println("写作得分为 : ", score)
//	return score
//}
//
//func (s *student) Selecting() int {
//	score := 42
//	fmt.Println("选择题得分为 : ", score)
//	return score
//}
//
//func (s *student) getTotalScore() int {
//	return s.Reading() + s.Listening() + s.Writing() + s.Listening()
//}
//
//func main() {
//	s := &student{
//		name : "小明",
//	}
//	score := s.getTotalScore()
//	fmt.Println("本次考试", s.name, "的得分为：", score)
//}