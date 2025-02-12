package main

import (
	"fmt"
	"os"

	"github.com/Qovop-Dev/My_Golang_Projects/All_Game_In_1/display"
	"github.com/Qovop-Dev/My_Golang_Projects/All_Game_In_1/game"
	"github.com/Qovop-Dev/My_Golang_Projects/All_Game_In_1/hangman"
	"github.com/Qovop-Dev/My_Golang_Projects/All_Game_In_1/helloyou"
	"github.com/Qovop-Dev/My_Golang_Projects/All_Game_In_1/input"
	"github.com/Qovop-Dev/My_Golang_Projects/All_Game_In_1/tictactoe"
)

func main() {

	//print project name
	display.DrawWelcome()

	//Welcome the user
	helloyou.Start()

	//Choose the game
	c, err := input.ReadChoice()
	if err != nil {
		fmt.Printf("Could not read from terminal: %v", err)
		os.Exit(1)
	}
	switch c {
	case "1":
		display.DrawWelcomeHangman()
		g, err := game.NewGame("Hangman")
		if err != nil {
			fmt.Printf("Could not open the game: %v\n", err)
			os.Exit(1)
		}
		hangman.Start(g)
	case "2", "3", "4":
		display.DrawWelcomeTicTacToe()
		g, err := game.NewGame("TicTacToe")
		if err != nil {
			fmt.Printf("Could not open the game: %v\n", err)
			os.Exit(1)
		}
		size := 0
		switch c {
		case "2":
			size = 3
		case "3":
			size = 4
		case "4":
			size = 5
		}
		tictactoe.Start(g, size)

	}

}
