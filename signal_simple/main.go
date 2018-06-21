package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

const (
	MAX_TO_QUIT = 3
)

func main() {
	fmt.Printf("Use (Ctrl + C) to capture the signal ... \n")
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	n := 0
	for sig := range c {
		if n == MAX_TO_QUIT {
			panic("got too many signals")
		}
		log.Printf("Recieved signal '%s' (%d times more to quit)", sig, MAX_TO_QUIT-n)
		n++
	}
}
