package tictactoe

import (
	"fmt"
)

func DrawWelcome() {

	//https://patorjk.com/software/taag
	//bubble font
	fmt.Println(`
---------------------------------------------------------------
   _   _   _   _   _   _   _   _   _   _   _  
  / \ / \ / \ / \ / \ / \ / \ / \ / \ / \ / \ 
 ( T | i | c | - | T | a | c | - | T | o | e )
  \_/ \_/ \_/ \_/ \_/ \_/ \_/ \_/ \_/ \_/ \_/ 

---------------------------------------------------------------`)
}

func DrawBoard(tg *TicTacToeGame) {

	alpha := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

	board := " "
	//alim en-tete
	for j := 1; j <= tg.Size; j++ {
		board += "   " + fmt.Sprint(j)
	}
	//alim lignes
	for i := 0; i < tg.Size; i++ {
		//valeur
		board += "\n" + alpha[i] + " | "
		for j := 0; j < tg.Size; j++ {
			board += tg.Move[i][j] + " | "
		}
		//séparateur
		board += "\n" + "  -"
		for j := 1; j < tg.Size; j++ {
			board += "---+"
		}
		board += "----"
	}
	fmt.Println(board)
}

func DrawState(tg *TicTacToeGame) {

	switch tg.State {
	case "won":
		fmt.Println("YOU WON!")

	case "lost":
		fmt.Println("You lost :( ! ")

	case "draw":
		fmt.Println("Nobody won...")
	}
}
