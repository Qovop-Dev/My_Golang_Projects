package dictionary

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

var words = make([]string, 0, 50)

func Load(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func PickWord() string {
	// Reset seed to be always different
	//rand.Seed(time.Now().Unix())
	rand.New(rand.NewSource(time.Now().Unix()))
	i := rand.Intn(len(words))
	return words[i]
}
