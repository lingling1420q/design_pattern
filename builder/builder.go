package main

import (
	"fmt"
)

type class interface {
	build() class
}


type classA struct {
	monitor string
	studentAssistant string
	chineseAssistant string
	mathAssistant string
	englishAssistant string
}

type classB struct {
	monitor string
	studentAssistant string
	chineseAssistant string
	mathAssistant string
	englishAssistant string
}

type ClassBuilder interface {
	buildMonitor(string) ClassBuilder
	buildStudentAssistant(string) ClassBuilder
	buildChineseAssistant(string) ClassBuilder
	buildMathAssistant(string) ClassBuilder
	buildEnglishAssistant(string) ClassBuilder
	build() class
}


func (c *classA) buildMonitor(s string) ClassBuilder {
	fmt.Println("自主报名...")
	fmt.Println("班级同学投票...")
	fmt.Println("统计票数，得出结果...")
	c.monitor = s
	return c
}

func (c *classB) buildMonitor(s string) ClassBuilder {
	fmt.Println("自主报名...")
	fmt.Println("科任老师和班主任共同指定...")
	c.monitor = s
	return c
}

func (c *classA) buildStudentAssistant(s string) ClassBuilder {
	fmt.Println("自主报名...")
	fmt.Println("班级同学投票...")
	fmt.Println("统计票数，得出结果...")
	c.studentAssistant = s
	return c
}

func (c *classA) buildChineseAssistant(s string) ClassBuilder {
	fmt.Println("自主报名...")
	fmt.Println("班级同学投票...")
	fmt.Println("统计票数，得出结果...")
	c.chineseAssistant = s
	return c
}

func (c *classA) buildMathAssistant(s string) ClassBuilder {
	fmt.Println("自主报名...")
	fmt.Println("班级同学投票...")
	fmt.Println("统计票数，得出结果...")
	c.mathAssistant = s
	return c
}

func (c *classA) buildEnglishAssistant(s string) ClassBuilder {
	fmt.Println("自主报名...")
	fmt.Println("班级同学投票...")
	fmt.Println("统计票数，得出结果...")
	c.englishAssistant = s
	return c
}

func (c *classA) build() class {
	return c
}

func (c *classB) buildStudentAssistant(s string) ClassBuilder {
	fmt.Println("自主报名...")
	fmt.Println("班级同学投票...")
	fmt.Println("统计票数，得出结果...")
	c.studentAssistant = s
	return c
}

func (c *classB) buildChineseAssistant(s string) ClassBuilder {
	fmt.Println("自主报名...")
	fmt.Println("班级同学投票...")
	fmt.Println("统计票数，得出结果...")
	c.chineseAssistant = s
	return c
}

func (c *classB) buildMathAssistant(s string) ClassBuilder {
	fmt.Println("自主报名...")
	fmt.Println("班级同学投票...")
	fmt.Println("统计票数，得出结果...")
	c.mathAssistant = s
	return c
}

func (c *classB) buildEnglishAssistant(s string) ClassBuilder {
	fmt.Println("自主报名...")
	fmt.Println("班级同学投票...")
	fmt.Println("统计票数，得出结果...")
	c.englishAssistant = s
	return c
}


func (c *classB) build() class {
	c.buildMonitor("小张").buildStudentAssistant("小吴")
	return c
}


func main() {
	c := &classA{}
	c.buildMonitor("小明").buildStudentAssistant("小红")

	fmt.Printf("A班级 ， 班长 ： %s, 学委 ： %s \n", c.monitor, c.studentAssistant)

	b := &classB{}

	fmt.Printf("B班级 ， 班长 ： %s, 学委 ： %s", b.monitor, b.studentAssistant)

}