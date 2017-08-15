package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	log.SetFlags(log.Ltime)

	var wg sync.WaitGroup
	semaphore := NewSemaphore(3)
	for i := 0; i < 9; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()
			defer semaphore()()

			// Resource intensive work goes here
			time.Sleep(time.Second)
			log.Println(i)
		}(i)
	}
	wg.Wait()
}

func NewSemaphore(max int) func() func() {
	c := make(chan struct{}, max)
	return func() func() {
		c <- struct{}{}
		return func() {
			<-c
		}
	}
}
