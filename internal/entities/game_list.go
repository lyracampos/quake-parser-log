package entities

import "fmt"

type GameList struct {
	Games      map[string]*Game
	gameNumber int
}

func NewGameList() *GameList {
	return &GameList{
		Games:      make(map[string]*Game, 0),
		gameNumber: 0,
	}
}

func (g *GameList) AddGame(game *Game) {
	g.gameNumber++

	gameName := g.formatGameName()
	g.Games[gameName] = game
}

func (g *GameList) formatGameName() string {
	return fmt.Sprintf("game_%02d", g.gameNumber)
}
