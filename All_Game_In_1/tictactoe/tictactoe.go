package tictactoe

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/Qovop-Dev/My_Golang_Projects/All_Game_In_1/display"
	"github.com/Qovop-Dev/My_Golang_Projects/All_Game_In_1/game"
	"github.com/Qovop-Dev/My_Golang_Projects/All_Game_In_1/input"
)

type ThisTicTacToeGame struct {
	*game.TicTacToeGame
}

func Start(g *game.Game, size int) {

	tg, err := NewTicTacToeGame(g, size)
	if err != nil {
		fmt.Printf("Could not create game: %v\n", err)
		os.Exit(1)
	}
	display.DrawBoard(tg.TicTacToeGame)

	for {
		switch tg.State {
		case "won", "lost", "draw":
			os.Exit(0)
		}

		l, c, err := input.ReadPosition(tg.TicTacToeGame)
		if err != nil {
			fmt.Printf("Could not read from terminal: %v", err)
			os.Exit(1)
		}

		tg.MakeAGuess(l, c)
	}
}

func NewTicTacToeGame(g *game.Game, size int) (*ThisTicTacToeGame, error) {

	move := make([][]string, size)

	for i := 0; i < size; i++ {
		move[i] = make([]string, size) // Initialize each sub-array
		for j := 0; j < size; j++ {
			move[i][j] = " "
		}
	}

	tg := &ThisTicTacToeGame{
		TicTacToeGame: &game.TicTacToeGame{
			Game: game.Game{
				Name:   g.Name,
				State:  g.State,
				Points: g.Points,
			},
			Move: move,
			Size: size,
		},
	}

	return tg, nil
}

func (tg *ThisTicTacToeGame) MakeAGuess(l, c int) {

	switch tg.State {
	case "won", "lost", "draw":
		return
	}

	//check if user has won
	if hasWon(tg.Move, tg.Size, "X") {
		tg.State = "won"
	}

	if tg.State != "won" {

		//check if game is over
		if isDraw(tg.Move) {
			tg.State = "draw"
		}

		if tg.State != "draw" {

			//qovop turn to play
			fmt.Println("\nQovop turn to Play 'O'")
			qovopPlay(tg.Move, tg.Size)
			display.DrawBoard(tg.TicTacToeGame)

			//check if qovop has won
			if hasWon(tg.Move, tg.Size, "O") {
				tg.State = "lost"
			}
		}
	}

	display.DrawStateTicTacToe(tg.TicTacToeGame)

}

func hasWon(moves [][]string, size int, player string) bool {

	//check win on all line
	for i := range moves {
		won := true
		for j := range moves[i] {
			if moves[i][j] != player {
				won = false
				break
			}
		}
		if won {
			return won
		}
	}

	//check win on all column
	for i := 0; i < size; i++ {
		won := true
		for j := range moves {
			if moves[j][i] != player {
				won = false
				break
			}
		}
		if won {
			return won
		}
	}

	//check win on 1st diag
	won := true
	for i := 0; i < size; i++ {
		if moves[i][i] != player {
			won = false
			break
		}
	}
	if won {
		return won
	}

	//check win on 1st diag
	won = true
	for i := 0; i < size; i++ {
		if moves[size-1-i][i] != player {
			won = false
			break
		}
	}
	if won {
		return won
	}

	return false
}

func isDraw(moves [][]string) bool {

	//check win on all line
	for i := range moves {
		for j := range moves[i] {
			if moves[i][j] == " " {
				return false
			}
		}
	}

	return true
}

func qovopPlay(moves [][]string, size int) {
	rand.New(rand.NewSource(time.Now().Unix()))
	for {
		i := rand.Intn(size)
		j := rand.Intn(size)
		if moves[i][j] == " " {
			moves[i][j] = "O"
			break
		}
	}
}
