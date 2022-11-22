package philosopher

import (
	"fmt"
	"sync"
	"time"
)

type Philosopher struct {
	Id int
}

func (p *Philosopher) Think() {
	fmt.Printf("Philosopher %d is thnking \n", p.Id)
}

func (p *Philosopher) Eat() {
	fmt.Printf("Philosopher %d is eating \n", p.Id)
	time.Sleep(1 * time.Second)
}

func (p *Philosopher) WantsToEat(wg *sync.WaitGroup, rf *Fork, lf *Fork) {
	fmt.Printf("Philosopher %d wants to eat \n", p.Id)
	select {
	case <-rf.Pick():
		fmt.Println("right first")
		<-lf.Pick()
		fmt.Println("left second")
	case <-lf.Pick():
		fmt.Println("left first")
		<-rf.Pick()
		fmt.Println("right second")
	}

	p.Eat()

	rf.Put() <- true
	lf.Put() <- true

	p.Think()
	wg.Done()
}

type Fork struct {
	Id     int
	Status chan bool
}

func NewFork(id int) *Fork {
	f := Fork{
		Id:     id,
		Status: make(chan bool, 1),
	}

	f.Status <- true
	return &f
}

func (f *Fork) Pick() <-chan bool {
	return f.Status
}

func (f *Fork) Put() chan<- bool {
	return f.Status
}
