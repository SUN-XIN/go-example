// blog: https://blog.golang.org/defer-panic-and-recover
package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	stopChan := make(chan bool)
	f1(stopChan)
	time.Sleep(3 * time.Second)
	stopChan <- true

	time.Sleep(time.Second)
	fmt.Println()
	log.Printf("Testing recover and panic now")
	f2(true)

	f2(false)
}

func f1(stopChan chan bool) {
	goFunc := func(stop chan bool) {
		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				log.Printf("I'm running")
			case <-stop:
				log.Printf("I stop")
			}
		}
	}

	defer func() {
		log.Printf("defer msg")
		go goFunc(stopChan)
	}()

	log.Printf("f1 msg")

	time.Sleep(time.Second)
}

func f2(runPanic bool) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("recover info: %+v", r)
		}
	}()

	log.Printf("f2 panic (%v)", runPanic)
	if runPanic {
		panic("panic in f2")
	}
	log.Printf("f2 end")
}
