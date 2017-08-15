package main

import (
	"log"
	"sync"
	"time"
)

func main() {
	log.SetFlags(log.Ltime)

	var wg sync.WaitGroup
	for i := 0; i < 9; i++ {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			// Resource intensive work goes here
			time.Sleep(time.Second)
			log.Println(i)
		}(i)
	}
	wg.Wait()
}
