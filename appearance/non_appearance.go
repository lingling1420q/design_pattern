package main

import "fmt"

type EnglishTest interface {
	Listening() int
	Reading() int
	Selecting() int
	Writing() int
}

type student struct {
	name string
}

func (s *student) listening(score int) int {
	fmt.Println(s.name , "听力得分为 : ", score)
	return score
}


func (s *student) selecting(score int) int {
	fmt.Println(s.name , "选择题得分为 : ", score)
	return score
}

func (s *student) reading(score int) int {
	fmt.Println(s.name , "阅读得分为 : ", score)
	return score
}

func (s *student) writing(score int) int {
	fmt.Println(s.name , "写作得分为 : ", score)
	if score > 20 {
		fmt.Println(s.name , "写作水平非常高")
	}
	return score
}

func (s *student) GetTotalScore() int {
	return s.listening(20) + s.selecting(42) + s.reading(40) + s.writing(25)
}

func main() {
	s := &student{
		name : "小明",
	}
	fmt.Println(s.name , "的总得分为 : ", s.GetTotalScore())

}