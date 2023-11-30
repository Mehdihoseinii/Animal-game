package gamerace

import (
	"fmt"
	"strings"
)

type Player interface {
	Move()
	Position() int
	Name() string
}
type Game struct {
	field_Length int
	players      []Player
}

func NewGame(field_Length int) *Game {
	return &Game{
		field_Length: field_Length,
		players:      make([]Player, 0),
	}
}

func (g *Game) Join(player Player) *Game {
	g.players = append(g.players, player)
	return g
}

func (g *Game) MovePlayers() {
	for _, player := range g.players {
		player.Move()
	}
}

func (g *Game) CheckWinner() Player {
	for _, player := range g.players {
		if player.Position() > g.field_Length {
			return player
		}
	}
	return nil
}

func (g *Game) Print() {
	fmt.Println("|" + strings.Repeat("-", g.field_Length+6) + "|")
	for _, player := range g.players {
		pos := player.Position()
		name := player.Name()
		row := "|" + strings.Repeat(" ", pos) + name + strings.Repeat(" ", g.field_Length-pos+1) + "|"
		fmt.Println(row)
	}

	fmt.Println("|" + strings.Repeat("-", g.field_Length+6) + "|")

}
