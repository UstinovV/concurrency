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

func (p *Philosopher) WantsToEat(wg *sync.WaitGroup, room *Room, rf *Fork, lf *Fork) {
	fmt.Printf("Philosopher %d wants to eat with %d and %d \n", p.Id, rf.Id, lf.Id)
	wait := time.NewTicker(100 * time.Millisecond)
	room.Enter() <- true
	fmt.Printf("Philosopher entered the room. Occupancy: %d \n", len(room.Occupancy))

	for {
		select {
		case <-rf.Pick():
			select {
			case <-lf.Pick():
				break
			case <-wait.C:
				rf.Put() <- true
				continue
			}

		case <-lf.Pick():
			select {
			case <-rf.Pick():
				break
			case <-wait.C:
				lf.Put() <- true
				continue
			}
		}
		break
	}
	p.Eat()

	select {
	case rf.Put() <- true:
		lf.Put() <- true
	case lf.Put() <- true:
		rf.Put() <- true
	}

	<-room.Exit()
	fmt.Printf("Philosopher left the room. Occupancy: %d \n", len(room.Occupancy))
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

type Room struct {
	Occupancy chan bool
	Forks     [5]Fork
}

func NewRoom() *Room {
	r := Room{
		Occupancy: make(chan bool, 4),
		Forks:     [5]Fork{},
	}

	return &r
}

func (r *Room) Enter() chan<- bool {
	return r.Occupancy
}

func (r *Room) Exit() <-chan bool {
	return r.Occupancy
}
