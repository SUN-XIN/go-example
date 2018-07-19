package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

const (
	RedirectMsg = "RedirectToAnotherURL"
)

func main() {
	clientRedirect := new(http.Client)
	clientRedirect.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return fmt.Errorf(RedirectMsg)
	}

	sendTo := "https://www.google.com"
	log.Printf("Test to send to %s, with clientRedirect", sendTo)
	resp, err := clientRedirect.Get(sendTo)
	checkError(err, resp)

	sendTo = "http://mydemo.effortless-lock-197213.appspot.com/redirect_to"
	log.Printf("Test to send to %s, with clientRedirect", sendTo)
	resp, err = clientRedirect.Get(sendTo)
	checkError(err, resp)

	client := new(http.Client)
	log.Printf("Test to send to %s, with client no Redirect", sendTo)
	resp, err = client.Get(sendTo)
	checkError(err, resp)
}

func handlerRedirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "http://www.google.com", 301)
}

func checkError(err error, resp *http.Response) {
	switch {
	case err == nil:
		log.Printf("No Redirect, code %d", resp.StatusCode)
	case strings.Contains(err.Error(), RedirectMsg):
		log.Printf("Redirect: msg %+v code %d", err, resp.StatusCode)
	default: // err no nil, not Redirect
		log.Printf("Failed send: %+v", err)
	}
}
