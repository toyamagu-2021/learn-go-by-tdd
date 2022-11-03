package main

import (
	"log"
	"net/http"

	poker "github.com/toyamagu-2021/learn-go-by-tdd/command-line"
)

const dbFileName = "game.db.json"

func main() {
	store, closeFunc, err := poker.FileSystemStoreFromFile(dbFileName)

	if err != nil {
		log.Fatal(err)

	}

	defer closeFunc()
	server := poker.NewPlayerServer(store)

	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listern on port 5000 %v", err)
	}
}
