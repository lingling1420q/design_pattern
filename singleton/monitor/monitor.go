package monitor

import (
	"fmt"
	"sync"
)

type Monitor interface {
	Help(string)
}

type monitor struct {
	name string
}

func (mo *monitor) Help(name string) {
	fmt.Printf(" %s 帮助了 %s \n", mo.name, name)
}

var mo *monitor

var mu sync.Mutex

var once sync.Once

func GetMonitor(name string) *monitor{

	once.Do(func() {
		mo = &monitor{
			name : name,
		}
		fmt.Println("1111111111111111111")
	})

	return mo
}

