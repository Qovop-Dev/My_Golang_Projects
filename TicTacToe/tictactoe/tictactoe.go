package tictactoe

import (
	"fmt"
	"math/rand"
	"time"
)

type TicTacToeGame struct {
	State string
	Move  [][]string
	Size  int
}

func New(size int) (*TicTacToeGame, error) {

	move := make([][]string, size)

	for i := 0; i < size; i++ {
		move[i] = make([]string, size) // Initialize each sub-array
		for j := 0; j < size; j++ {
			move[i][j] = " "
		}
	}

	tg := &TicTacToeGame{
		State: "",
		Move:  move,
		Size:  size,
	}

	DrawBoard(tg)

	return tg, nil
}

func (tg *TicTacToeGame) MakeAGuess() {

	switch tg.State {
	case "won", "lost", "draw":
		return
	}

	//check if user has won
	if hasWon(tg.Move, tg.Size, "X") {
		tg.State = "won"
		DrawState(tg)
		return
	}

	//check if game is over
	if isDraw(tg.Move) {
		tg.State = "draw"
		DrawState(tg)
		return
	}

	//qovop turn to play
	fmt.Println("\nQovop turn to Play 'O'")
	qovopPlay(tg.Move, tg.Size)
	DrawBoard(tg)

	//check if qovop has won
	if hasWon(tg.Move, tg.Size, "O") {
		tg.State = "lost"
		DrawState(tg)
		return
	}

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
