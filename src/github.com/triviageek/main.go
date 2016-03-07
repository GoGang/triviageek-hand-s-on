package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"

	"github.com/triviageek/game"
)

/*
 - MAIN FILE -
 Initialize game package and bind handler to default path /
*/
func main() {
	fmt.Println("Application starting...")

	game.Init()

	http.Handle("/", websocket.Handler(GameHandler))

	err := http.ListenAndServe(":9001", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}

}

func GameHandler(ws *websocket.Conn) {
	defer func(ws *websocket.Conn) {
		if err := recover(); err != nil {
			ws.Close()
		}
	}(ws)

	fmt.Println("Incoming connection from", ws.RemoteAddr())

	receivedBytes := make([]byte, 100)
	n, err := ws.Read(receivedBytes)
	if err != nil {
		fmt.Println("Not a websocket request, close connection", err)
		panic(err)
	}
	p := &game.Player{}
	if err := json.Unmarshal(receivedBytes[:n], p); err != nil {
		fmt.Println("Bad request, object is not a player. Close connection", err)
		panic(err)
	}
	p.Ws = ws

	fmt.Println("New player joining the game :", *p)

	p.JoinAGame()
	p.HandleEvents()
}
