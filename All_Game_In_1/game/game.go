package game

type User struct {
	Name string
}

type Game struct {
	Name   string
	State  string
	Points float32
}

type HangmanGame struct {
	Game
	Letters      []string
	FoundLetters []string
	UsedLetters  []string
	TurnsLeft    int
}

type TicTacToeGame struct {
	Game
	Move [][]string
	Size int
}

func NewGame(gameName string) (*Game, error) {

	g := &Game{
		Name:   gameName,
		State:  "",
		Points: 0.00,
	}
	return g, nil
}
