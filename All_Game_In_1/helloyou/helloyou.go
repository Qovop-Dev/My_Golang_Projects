package helloyou

import (
	"fmt"
	"os"
	"strings"

	"github.com/Qovop-Dev/My_Golang_Projects/All_Game_In_1/input"
)

func Start() {

	n, err := input.AskName()
	if err != nil {
		fmt.Printf("Could not read from terminal: %v", err)
		os.Exit(1)
	}
	name := strings.ToUpper(n.Name[:1]) + strings.ToLower(n.Name[1:])
	fmt.Printf("\nWelcome to you %v!\n\n", name)

}
