package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func main() {
	// Test 1
	blockedChan := make(chan bool)
	go func() {
		for {
			<-blockedChan
			log.Printf("blockedChan gets 1 msg, pending 1s...")
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 3; i++ {
		log.Printf("Sending to blockedChan ...")
		blockedChan <- true
	}
	fmt.Println()

	// Test 2
	BufferChan := make(chan bool, 3)
	go func() {
		for {
			<-BufferChan
			log.Printf("BufferChan gets 1 msg, pending 1s...")
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 3; i++ {
		log.Printf("Sending to BufferChan ...")
		BufferChan <- true
	}
	time.Sleep(3 * time.Second)
	fmt.Println()

	// Test 3
	closeByProducerChan := make(chan bool)
	go func() {
		var ok bool
		for {
			_, ok = <-closeByProducerChan
			if !ok {
				log.Printf("closeByProducerChan chan is closed -> over")
				return
			}
			log.Printf("closeByProducerChan gets 1 msg, pending 1s...")
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < 3; i++ {
		log.Printf("Sending to closeByProducerChan ...")
		closeByProducerChan <- true
	}
	close(closeByProducerChan)
	time.Sleep(time.Second)
	fmt.Println()

	// Test 4
	closeByConsumer := make(chan bool)
	go func() {
		for i := 0; i < 3; i++ {
			<-closeByConsumer
			log.Printf("closeByConsumer gets 1 msg, pending 1s...")
			time.Sleep(time.Second)
		}
		log.Printf("I'm consumer, I close chan")
		close(closeByConsumer)
	}()

	defer func() {
		if r := recover(); r != nil {
			if strings.Contains(fmt.Sprintf("%v", r), "send on closed channel") {
				log.Printf("I'm producter, closeByConsumer is closed by consumer")
			} else {
				log.Printf("recover info: %+v", r)
			}
		}
	}()
	for i := 0; i < 5; i++ {
		log.Printf("Sending to closeByConsumer ...")
		closeByConsumer <- true
	}

	time.Sleep(time.Minute)
}
