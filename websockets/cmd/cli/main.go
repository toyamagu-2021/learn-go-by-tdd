package main

import (
	"fmt"
	"log"
	"os"

	poker "github.com/toyamagu-2021/learn-go-by-tdd/ws"
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
	game := poker.NewTexasHoldem(poker.BlindAlerterFunc(poker.Alerter), store)
	poker.NewCLI(os.Stdin, os.Stdout, game).PlayPoker()
}
