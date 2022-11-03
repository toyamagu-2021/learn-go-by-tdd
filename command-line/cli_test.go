package poker_test

import (
	"strings"
	"testing"

	poker "github.com/toyamagu-2021/learn-go-by-tdd/command-line"
)

func TestCLI(t *testing.T) {
	t.Run("record Taro from user input", func(t *testing.T) {
		in := strings.NewReader("Taro wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Taro")
	})

	t.Run("record Takeshi from user input", func(t *testing.T) {
		in := strings.NewReader("Takeshi wins\n")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWin(t, playerStore, "Takeshi")
	})

}
