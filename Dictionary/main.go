package main

import (
	"dictionary/dictionary"
	"flag"
	"fmt"
	"os"
)

func main() {
	d, err := dictionary.New("./badger")
	handleErr(err)
	defer d.Close()

	action := flag.String("action", "list", "Type of action to perform on the dictionary")
	flag.Parse()

	switch *action {
	case "add":
		actionAdd(d, flag.Args())
		//d.Add("Egypt", "A wonderful country")
	case "get":
		actionGet(d, flag.Args())
	case "list":
		actionList(d)
	case "del":
		actionDelete(d, flag.Args())
	default:
		fmt.Println("Action name not valid")
	}
	/*// Run garbage collection
	err = d.RunGC()
	handleErr(err)*/
}

func actionList(d *dictionary.Dictionary) {
	words, entries, err := d.List()
	handleErr(err)
	fmt.Println("Dictionary Content:")
	for _, word := range words {
		fmt.Println(entries[word])
	}
}

func actionAdd(d *dictionary.Dictionary, args []string) {
	word := args[0]
	definition := args[1]
	err := d.Add(word, definition)
	handleErr(err)
	fmt.Printf("%v added to the dictionary\n", word)
}

func actionGet(d *dictionary.Dictionary, args []string) {
	word := args[0]
	entry, err := d.Get(word)
	handleErr(err)
	fmt.Println(entry)
}

func actionDelete(d *dictionary.Dictionary, args []string) {
	word := args[0]
	err := d.Del(word)
	handleErr(err)
	fmt.Printf("'%v' was removed from the dictionary\n", word)
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("Dictionary error: %v\n", err)
		os.Exit(1)
	}
}
