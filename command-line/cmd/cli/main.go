package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/toyamagu-2021/learn-go-by-tdd/command-line"
)

const dbFileName = "game.db.json"

func main() {

	store, closeFunc, err := poker.FileSystemStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)

	}

	defer closeFunc()

	fmt.Println("Let's play poker")
	fmt.Println(("Type {Name} wins to record a win"))
	poker.NewCLI(store, os.Stdin).PlayPoker()
}
