package main

import (
	"log"
	"time"
)

func main() {
	defer func() {
		log.Printf("in defer func, pending...")
		time.Sleep(2 * time.Second)
		log.Printf("defer func is over")
	}()

	log.Printf("out of defer")

	time.Sleep(2 * time.Second)
	log.Printf("panic soon ...\n(defer before or after panic ???)")
	panic("STOP")
}
