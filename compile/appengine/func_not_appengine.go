// +build !appengine

package main

import (
	"log"
	"net/http"
)

func Send() {
	client := new(http.Client)
	_, err := client.Get("http://www.google.com/")
	if err != nil {
		log.Printf("Failed Get: %+v", err)
		return
	}

	log.Printf("Get ok")
}
