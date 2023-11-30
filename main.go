package main

import (
	"animalrace/gamerace"
	"animalrace/tools"
	"fmt"
	"time"
)

func main() {
	//macke plaeyer
	dog := gamerace.NewDog("Alex")
	cat := gamerace.NewCat("Bincky")
	owl := gamerace.NewOwl("Hanna")

	//macke game

	game := gamerace.NewGame(50)

	//join pleayer to the game
	game.Join(cat).Join(dog).Join(owl)
	//loop till ther is a winner
	var winner gamerace.Player
	for winner == nil {
		//Move the pleayer
		game.MovePlayers()
		//checke winner

		winner = game.CheckWinner()
		//Print Game
		game.Print()
		time.Sleep(time.Millisecond * 250)
		tools.ClearTerminal()
	}

	//Print the winner

	fmt.Printf("%s won the game!", winner.Name())
}
