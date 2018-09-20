package main

import "log"

func main() {
	log.Printf("%v", Sum(1, 2))
	//log.Printf("%v", Sum(1.1, 2.2)) // must compile with tags "mytag"
}
