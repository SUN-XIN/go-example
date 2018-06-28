package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	ws "github.com/gorilla/websocket"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var (
	isClosed      = false
	receivePerson *Person
)

func main() {
	http.HandleFunc("/testws", connHandler)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func connHandler(w http.ResponseWriter, r *http.Request) {
	upgrader := ws.Upgrader{
		//ReadBufferSize:  4096,
		//WriteBufferSize: 4096,
		CheckOrigin: func(r *http.Request) bool {
			// extract the given key from url
			key := r.URL.Query().Get("key")
			if key == "" {
				log.Printf("Error KeyNotFound: %s", r.URL.String())
				return false
			}

			// TODO: check key

			log.Printf("verify key (%s) OK", key)

			return true
		},
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Printf("Error upgrader.Upgrade: %+v \n", err)
		return
	}

	isClosed = false
	go sendMessage(conn)
	go receiveMessage(conn)
}

func receiveMessage(conn *ws.Conn) {
	for !isClosed {
		// wait the new message from this client
		_, message, err := conn.ReadMessage()
		if err != nil {
			if connIsClosed(err) {
				log.Printf("conn is closed")
				return
			}
			log.Printf("Error ReadMessage: %+v", err)
			continue
		}

		log.Printf("get msg: %s", message)

		var p Person
		err = json.Unmarshal(message, &p)
		if err != nil {
			log.Printf("Failed Unmarshal to Person: %+v", err)
		}
		receivePerson = &p
	}
}

func sendMessage(conn *ws.Conn) {
	for {
		if receivePerson == nil {
			time.Sleep(time.Second)
			continue
		}

		break
	}
	for i := 0; i < 5; i++ {
		msg := []string{
			receivePerson.Name,
			strconv.Itoa(receivePerson.Age),
			strconv.Itoa(i),
		}
		conn.WriteMessage(ws.TextMessage, []byte(strings.Join(msg, ",")))
		log.Printf("msg %d is sent", i)
		time.Sleep(time.Second)
	}

	isClosed = true
	conn.Close()
	log.Printf("close conn")
}

func connIsClosed(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), "use of closed network connection")
}
