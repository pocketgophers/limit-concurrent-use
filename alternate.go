package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	log.SetFlags(log.Ltime)

	var wg sync.WaitGroup
	semaphore := NewSemaphore(4)
	for i := 0; i < 20; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			semaphore.Acquire()
			defer semaphore.Release()

			// Resource intensive work here
			time.Sleep(time.Second)
			log.Println(i)
		}(i)
	}
	wg.Wait()
}

type Semaphore chan struct{}

func NewSemaphore(max int) Semaphore {
	return make(chan struct{}, max)
}

func (s Semaphore) Acquire() {
	s <- struct{}{}
}

func (s Semaphore) Release() {
	<-s
}
