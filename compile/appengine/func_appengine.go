// +build appengine

package main

import (
	"context"
	"log"

	"google.golang.org/appengine/urlfetch"
)

func Send(ctx context.Context) {
	clientApp := urlfetch.Client(ctx)
	_, err := clientApp.Get("http://www.google.com/")
	if err != nil {
		log.Printf("Failed Get: %+v", err)
		return
	}

	log.Printf("Get ok")
}
