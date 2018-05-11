package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const (
	proxyIP   = "167.99.107.202"
	proxyPort = "3128"

	sendToURL = "http://mydemo.effortless-lock-197213.appspot.com/my_ip"
)

func main() {
	fmt.Printf("sending to url: %s \n", sendToURL)

	proxy := func(_ *http.Request) (*url.URL, error) {
		return url.Parse(fmt.Sprintf("http://%s:%s", proxyIP, proxyPort))
	}
	transport := &http.Transport{Proxy: proxy}

	c := &http.Client{Transport: transport}

	req, err := http.NewRequest("GET", sendToURL, nil)
	if err != nil {
		fmt.Printf("Failed NewRequest: %+v", err)
		return
	}

	res, err := c.Do(req)
	if err != nil {
		fmt.Printf("Failed send request: %+v", err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Failed ReadAll: %+v", err)
		return
	}

	fmt.Printf("%s \n", body)
}
