package main

import (
	"fmt"
	"sync"
	"time"

	ph "github.com/UstinovV/concurrency/the-dining-philosophers/philosopher"
)

const (
	N = 1
	P = 5
)

func main() {
	start := time.Now()
	philosophers := [5]ph.Philosopher{}
	room := ph.NewRoom()
	for i := 0; i < P; i++ {
		room.Forks[i] = *ph.NewFork(i)
		philosophers[i] = ph.Philosopher{Id: i}
	}

	wg := &sync.WaitGroup{}
	wg.Add(P * N)

	for i := 0; i < N; i++ {
		for j := 0; j < P; j++ {
			go philosophers[j].WantsToEat(wg, room, &room.Forks[j], &room.Forks[(j+1)%5])
		}
	}
	wg.Wait()

	duration := time.Since(start)
	fmt.Println(duration)
}
