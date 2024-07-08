package helloyou

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type User struct {
	Name string
}

var reader = bufio.NewReader(os.Stdin)

func AskName() (u User, err error) {
	valid := false
	for !valid {
		fmt.Print("What is your Name ? ")
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

func (u User) HelloUser() {
	name := strings.ToUpper(u.Name[:1]) + strings.ToLower(u.Name[1:])
	fmt.Printf("\nHello my new Friend %v!\nWelcome to my GitHub.\n\n", name)
}
