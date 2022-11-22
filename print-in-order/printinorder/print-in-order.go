package printinorder

import (
	"fmt"
)

type Foo struct {
	fs chan bool
	st chan bool
	tm chan bool
}

func (f *Foo) first() {
	fmt.Print("first")
	f.fs <- true

}

func (f *Foo) second() {
	<-f.fs
	fmt.Print("second")
	f.st <- true
}

func (f *Foo) third() {
	<-f.st
	fmt.Println("third")
	f.tm <- true
}

func PrintInOrder(order [3]int) {
	f := &Foo{
		fs: make(chan bool),
		st: make(chan bool),
		tm: make(chan bool),
	}

	for _, i := range order {
		fmt.Println(i)
		switch i {
		case 1:
			go func(f *Foo) {
				f.first()
			}(f)
		case 2:
			go func(f *Foo) {
				f.second()
			}(f)
		case 3:
			go func(f *Foo) {
				f.third()
			}(f)
		}
	}

	<-f.tm
}
