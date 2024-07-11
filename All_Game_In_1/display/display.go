package display

import (
	"fmt"

	"github.com/Qovop-Dev/My_Golang_Projects/All_Game_In_1/game"
)

func DrawWelcome() {
	//https://patorjk.com/software/taag
	//bubble font
	fmt.Println(`
---------------------------------------------------------------
   _   _   _     _   _   _   _     _   _     _  
  / \ / \ / \   / \ / \ / \ / \   / \ / \   / \ 
 ( A | l | l ) ( G | a | m | e ) ( I | n ) ( 1 )
  \_/ \_/ \_/   \_/ \_/ \_/ \_/   \_/ \_/   \_/  
 
---------------------------------------------------------------
 `)
}

func DrawWelcomeHangman() {
	//https://patorjk.com/software/taag
	//bubble font
	fmt.Println(`
---------------------------------------------------------------
   _   _   _   _   _   _   _  
  / \ / \ / \ / \ / \ / \ / \ 
 ( H | a | n | g | m | a | n )
  \_/ \_/ \_/ \_/ \_/ \_/ \_/ 

---------------------------------------------------------------`)
}

func Draw(hg *game.HangmanGame, guess string) {
	drawTurns(hg.TurnsLeft)
	drawState(hg, guess)
	fmt.Println()
}

func drawTurns(l int) {
	var draw string
	switch l {
	case 0:
		draw = `
    ____
   |    |      
   |    o      
   |   /|\     
   |    |
   |   / \
  _|_
 |   |______
 |          |
 |__________|
		`
	case 1:
		draw = `
    ____
   |    |      
   |    o      
   |   /|\     
   |    |
   |    
  _|_
 |   |______
 |          |
 |__________|
		`
	case 2:
		draw = `
    ____
   |    |      
   |    o      
   |    |
   |    |
   |     
  _|_
 |   |______
 |          |
 |__________|
		`
	case 3:
		draw = `
    ____
   |    |      
   |    o      
   |        
   |   
   |   
  _|_
 |   |______
 |          |
 |__________|
		`
	case 4:
		draw = `
    ____
   |    |      
   |      
   |      
   |  
   |  
  _|_
 |   |______
 |          |
 |__________|
		`
	case 5:
		draw = `
    ____
   |        
   |        
   |        
   |   
   |   
  _|_
 |   |______
 |          |
 |__________|
		`
	case 6:
		draw = `
    
   |     
   |     
   |     
   |
   |
  _|_
 |   |______
 |          |
 |__________|
		`
	case 7:
		draw = `
  _ _
 |   |______
 |          |
 |__________|
		`
	case 8:
		draw = `

		`
	}
	fmt.Println(draw)
}

func drawLetters(g []string) {
	for _, c := range g {
		fmt.Printf("%v ", c)
	}
	fmt.Println()
}

func drawState(g *game.HangmanGame, guess string) {
	fmt.Print("Guessed: ")
	drawLetters(g.FoundLetters)

	fmt.Print("Used: ")
	drawLetters(g.UsedLetters)

	switch g.State {
	case "goodGuess":
		fmt.Println("Good guess!")
	case "alreadyGuessed":
		fmt.Printf("Letter '%s' was already used.\n", guess)
	case "badGuess":
		fmt.Printf("Bad guess, '%s' is not in the word.\n", guess)
	case "lost":
		fmt.Print("You lost :(! The word was: ")
		drawLetters(g.Letters)
	case "won":
		fmt.Print("\nYOU WON! The word was: ")
		drawLetters(g.Letters)
		if g.Points < 1 {
			g.Points = 1
		}
		fmt.Printf("\nYou got %v points!", g.Points)
	}
	fmt.Println()
}
