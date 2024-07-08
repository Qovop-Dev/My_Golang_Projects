package hangman

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
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
