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

var (
	table        [P]ph.Fork
	philosophers [P]ph.Philosopher
)

func main() {
	start := time.Now()

	for i := 0; i < P; i++ {
		table[i] = *ph.NewFork(i)
		philosophers[i] = ph.Philosopher{Id: i}
	}

	wg := &sync.WaitGroup{}
	wg.Add(P * N)

	for i := 0; i < N; i++ {
		for j := 0; j < P; j++ {
			go philosophers[j].WantsToEat(wg, &table[j], &table[(j+1)%5])
		}
	}
	wg.Wait()

	duration := time.Since(start)
	fmt.Println(duration)
}
