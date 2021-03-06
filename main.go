package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ahmedMunna1767/go-vue-socket-chat/pkg/websocket"
)

func serveWs(pool *websocket.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := websocket.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websocket.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Server Pinged")
	w.Write([]byte("This is Hello World From The Chat Server"))
}

func setupRoutes() {
	pool := websocket.NewPool()
	go pool.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(pool, w, r)
	})

	http.HandleFunc("/", hello)

}

func main() {
	fmt.Println("Go Chat App with TLS Config")
	setupRoutes()
	log.Fatal(http.ListenAndServe(":443", nil))
}
