package main

import (
	"log"
	"time"
)

func main() {
	log.SetFlags(log.Ltime)

	for i := 0; i < 9; i++ {
		// Resource intensive work goes here
		time.Sleep(time.Second)
		log.Println(i)
	}
}
