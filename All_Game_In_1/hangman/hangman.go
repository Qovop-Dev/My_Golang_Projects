package hangman

import (
	"fmt"
	"os"
	"strings"

	"github.com/Qovop-Dev/My_Golang_Projects/All_Game_In_1/dictionary"
	"github.com/Qovop-Dev/My_Golang_Projects/All_Game_In_1/display"
	"github.com/Qovop-Dev/My_Golang_Projects/All_Game_In_1/game"
	"github.com/Qovop-Dev/My_Golang_Projects/All_Game_In_1/input"
)

type ThisHangmanGame struct {
	*game.HangmanGame
}

func Start(g *game.Game) {
	err := dictionary.Load("dictionary/words.txt")
	if err != nil {
		fmt.Printf("Could not load dictionary: %v\n", err)
		os.Exit(1)
	}

	hg, err := NewHangmanGame(g, 8, dictionary.PickWord())
	if err != nil {
		fmt.Printf("Could not create game: %v\n", err)
		os.Exit(1)
	}

	guess := ""
	for {
		display.Draw(hg.HangmanGame, guess)

		switch hg.State {
		case "won", "lost":
			os.Exit(0)
		}

		l, err := input.ReadGuess()
		if err != nil {
			fmt.Printf("Could not read from terminal: %v", err)
			os.Exit(1)
		}
		guess = l

		hg.MakeAGuess(guess)
	}
}

func NewHangmanGame(g *game.Game, turns int, word string) (*ThisHangmanGame, error) {
	if len(word) < 3 {
		return nil, fmt.Errorf("word '%s' must be at least 3 characters. got=%v", word, len(word))
	}

	letters := strings.Split(strings.ToUpper(word), "")
	found := make([]string, len(letters))
	for i := 0; i < len(letters); i++ {
		found[i] = "_"
	}

	hg := &ThisHangmanGame{
		HangmanGame: &game.HangmanGame{
			Game: game.Game{
				Name:   g.Name,
				State:  g.State,
				Points: 10.0,
			},
			Letters:      letters,
			FoundLetters: found,
			TurnsLeft:    turns,
		},
	}
	fmt.Printf("Name%v\n", hg.Name)
	fmt.Printf("State%v\n", hg.State)
	fmt.Printf("Points%v\n", hg.Points)

	return hg, nil
}

func (hg *ThisHangmanGame) MakeAGuess(guess string) {
	guess = strings.ToUpper(guess)

	switch hg.State {
	case "won", "lost":
		return
	}

	if letterInWord(guess, hg.UsedLetters) {
		hg.State = "alreadyGuessed"
		hg.Points -= 0.5
	} else if letterInWord(guess, hg.Letters) {
		hg.State = "goodGuess"
		hg.RevealLetter(guess)

		if hasWon(hg.Letters, hg.FoundLetters) {
			hg.State = "won"
		}
	} else {
		hg.LoseTurn(guess)
		hg.Points -= 1

		if hg.TurnsLeft <= 0 {
			hg.State = "lost"
		}
	}
}

func hasWon(letters []string, foundLetters []string) bool {
	for i := range letters {
		if letters[i] != foundLetters[i] {
			return false
		}
	}
	return true
}

func (hg *ThisHangmanGame) RevealLetter(guess string) {
	hg.UsedLetters = append(hg.UsedLetters, guess)
	for i, l := range hg.Letters {
		if l == guess {
			hg.FoundLetters[i] = guess
		}
	}
}

func (hg *ThisHangmanGame) LoseTurn(guess string) {
	hg.TurnsLeft--
	hg.UsedLetters = append(hg.UsedLetters, guess)
}

func letterInWord(guess string, letters []string) bool {
	for _, l := range letters {
		if l == guess {
			return true
		}
	}
	return false
}
