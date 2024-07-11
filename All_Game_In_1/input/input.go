package input

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/Qovop-Dev/My_Golang_Projects/All_Game_In_1/game"
)

var reader = bufio.NewReader(os.Stdin)

func ReadGuess() (guess string, err error) {
	valid := false
	for !valid {
		fmt.Print("What is your letter ? ")
		guess, err = reader.ReadString('\n')
		if err != nil {
			return guess, err
		}
		guess = strings.TrimSpace(guess)
		if len(guess) != 1 {
			fmt.Printf("Invalid letter input (letter =%v, len=%v), only one letter.\n", guess, len(guess))
			continue
		}

		re := regexp.MustCompile(`[^a-zA-Z]`)
		if re.MatchString(guess) {
			fmt.Printf("Invalid letter input (letter =%v), must be alphabetic.\n", guess)
			continue
		}
		valid = true
	}
	return guess, nil
}

func AskName() (u game.User, err error) {
	valid := false
	for !valid {
		fmt.Print("Hello new Player, what is your Name ? ")
		u.Name, err = reader.ReadString('\n')
		if err != nil {
			return u, err
		}
		u.Name = strings.TrimSpace(u.Name)
		if len(u.Name) <= 1 {
			fmt.Printf("Invalid name input (letter =%v, len=%v), must be more than one character.\n", u.Name, len(u.Name))
			continue
		}

		re := regexp.MustCompile(`[^a-zA-Z- ]`)
		if re.MatchString(u.Name) {
			fmt.Printf("Invalid name input (letter =%v), must be alphabetic.\n", u.Name)
			continue
		}
		valid = true
	}
	return u, nil
}

func ReadChoice() (choice string, err error) {
	valid := false
	for !valid {
		fmt.Print("Choose your game:\n\n\t1) Hangman\n\n\t2) Tic-Tac-Toe\n\nYour Choice: ")
		choice, err = reader.ReadString('\n')
		if err != nil {
			return choice, err
		}
		choice = strings.TrimSpace(choice)
		if len(choice) != 1 {
			fmt.Printf("Invalid choice input (letter =%v, len=%v), only one figure.\n", choice, len(choice))
			continue
		}

		re := regexp.MustCompile(`[^1-2]`)
		if re.MatchString(choice) {
			fmt.Printf("\nInvalid choice input (choice = %v).\n\n", choice)
			continue
		}
		valid = true
	}
	return choice, nil
}
