package main

import (
	"fmt"
	"os"

	"github.com/Qovop-Dev/My_Golang_Projects/Hello_You/helloyou"
)

func main() {

	n, err := helloyou.AskName()
	if err != nil {
		fmt.Printf("Could not read from terminal: %v", err)
		os.Exit(1)
	}
	n.HelloUser()
}
