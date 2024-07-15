package main

import (
	"fmt"
	"os"

	"github.com/Qovop-Dev/My_Golang_Projects/TicTacToe/tictactoe"
)

func main() {
	tictactoe.DrawWelcome()
	choice, err := tictactoe.ReadChoice()
	if err != nil {
		fmt.Printf("Could not read your choice: %v\n", err)
		os.Exit(1)
	}
	size := 0
	switch choice {
	case "1":
		size = 3
	case "2":
		size = 4
	case "3":
		size = 5
	}
	tg, err := tictactoe.New(size)
	if err != nil {
		fmt.Printf("Could not create game: %v\n", err)
		os.Exit(1)
	}

	for {
		switch tg.State {
		case "won", "lost", "draw":
			os.Exit(0)
		}

		l, c, err := tictactoe.ReadGuess(tg)
		if err != nil {
			fmt.Printf("Could not read from terminal: %v", err)
			os.Exit(1)
		}

		tg.MakeAGuess(l, c)
	}

}
