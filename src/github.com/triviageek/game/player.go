package game

import (
	"golang.org/x/net/websocket"
)

type Player struct {
	Pseudo string          `json:"pseudo"`
	Score  int             `json:"score"`
	Ws     *websocket.Conn `json:"-"`
}

type Response struct {
	Step  int    `json:"step"`
	Value string `json:"value"`
}

func (p *Player) JoinAGame() {

}

func (p *Player) HandleEvents() {

}
