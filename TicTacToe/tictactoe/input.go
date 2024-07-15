package tictactoe

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"unicode"
)

var reader = bufio.NewReader(os.Stdin)

func ReadChoice() (choice string, err error) {
	valid := false
	for !valid {
		fmt.Print("Choose your game:\n\n\t1) Tic-Tac-Toe (3x3)\n\n\t2) Tic-Tac-Toe (4x4)\n\n\t3) Tic-Tac-Toe (5x5)\n\nYour Choice: ")
		choice, err = reader.ReadString('\n')
		if err != nil {
			return choice, err
		}
		choice = strings.TrimSpace(choice)
		if len(choice) != 1 {
			fmt.Printf("Invalid choice input (letter =%v, len=%v), only one figure.\n", choice, len(choice))
			continue
		}

		re := regexp.MustCompile(`[^1-3]`)
		if re.MatchString(choice) {
			fmt.Printf("\nInvalid choice input (choice = %v).\n\n", choice)
			continue
		}
		valid = true
	}
	return choice, nil
}

func ReadGuess(tg *TicTacToeGame) (l, c int, err error) {
	alpha := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	valid := false
	for !valid {
		fmt.Print("What is your position ? ")
		guess, err := reader.ReadString('\n')
		if err != nil {
			return l, c, err
		}
		guess = strings.TrimSpace(guess)
		if len(guess) != 2 {
			fmt.Printf("Invalid letter input (letter = %v, len = %v), only two letter.\n", guess, len(guess))
			continue
		}

		zone1 := "[^a-" + strings.ToLower(alpha[tg.Size-1]) + "A-" + alpha[tg.Size-1] + "]"
		re := regexp.MustCompile(zone1)
		if re.MatchString(string(guess[0])) {
			fmt.Printf("Invalid first letter input (letter = %v).\n", guess)
			continue
		}

		zone2 := "[^1-" + fmt.Sprint(tg.Size) + "]"
		re = regexp.MustCompile(zone2)
		if re.MatchString(string(guess[1])) {
			fmt.Printf("Invalid second letter input (letter = %v).\n", guess)
			continue
		}

		l := int(unicode.ToUpper(rune(guess[0])) - 'A')
		c := int(guess[1]-'0') - 1
		if tg.Move[l][c] != " " {
			fmt.Printf("This position is already used (Position = %v).\n", strings.ToUpper(guess))
			continue
		}

		valid = true
		tg.Move[l][c] = "X"
		DrawBoard(tg)

	}
	return l, c, nil
}
