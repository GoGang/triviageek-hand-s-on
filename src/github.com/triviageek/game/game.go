package game

import (
	"time"
)

const (
	NumOfQuestionsPerGame = 12
	QuestionPeriod        = 20
)

var runningGames []*Game

type Game struct {
	Name      string    `json:"name"`
	StartTime time.Time `json:"startTime"`
	Step      int       `json:"step"`
}

type Result struct {
	Players []*Player `json:"players"`
}

type Question struct {
	Step        int      `json:"step"`
	Smell       Smell    `json:"smell"`
	Suggestions []string `json:"suggestions"`
}

func (g *Game) start() {
	// Send a question every 20 sec

}
