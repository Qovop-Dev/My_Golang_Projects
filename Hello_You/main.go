package main

import (
	"fmt"
	"helloyou/helloyou"
	"os"
)

func main() {

	n, err := helloyou.AskName()
	if err != nil {
		fmt.Printf("Could not read from terminal: %v", err)
		os.Exit(1)
	}
	n.HelloUser()
}
