package main

import "log"

type Person struct {
	Name string
	Age  int8
}

func main() {
	pNew := new(Person)
	log.Printf("init pointer with new -> (%v)", pNew)

	var ppVar *Person
	log.Printf("init pointer with var -> (%v)", ppVar)

	pp := &Person{}
	log.Printf("init pointer without val -> (%v)", pp)

	var p Person
	log.Printf("init with var -> (%v)", p)

	pers := Person{}
	log.Printf("init without val -> (%v)", pers)
}
