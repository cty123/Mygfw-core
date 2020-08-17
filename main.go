package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:    4096,
	WriteBufferSize:   4096,
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Got incomming connection.")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Upgrade:", err)
		return
	}
	defer conn.Close()

	for {
		_, b, err := conn.ReadMessage()
		if err != nil {
			if err != io.EOF {
				log.Println("NextReader:", err)
			}
			return
		}
		log.Println(b)
	}
}

var addr = flag.String("addr", ":9000", "http service address")

func main() {
	flag.Parse()
	http.HandleFunc("/test", handler)
	fmt.Println("Server started on", *addr)
	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
