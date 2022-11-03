package poker_test

import (
	"bytes"
	"fmt"
	"io"
	"strings"
	"testing"
	"time"

	poker "github.com/toyamagu-2021/learn-go-by-tdd/ws"
)

var dummyBlindAlerter = &poker.SpyBlindAlerter{}
var dummyPlayerStore = &poker.StubPlayerStore{}

func TestCLI(t *testing.T) {
	t.Run("record Taro from user input", func(t *testing.T) {
		in := strings.NewReader("5\nChris wins")
		out := &bytes.Buffer{}
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, out, game)
		cli.PlayPoker()

		poker.AssertMessagesSentToUser(t, out, poker.PlayerPrompt)
		poker.AsserGameStartedWith(t, game, 5)
		poker.AsserGameFinishedWith(t, game, "Chris")
	})

	t.Run("it prints an error when a non numeric value", func(t *testing.T) {
		in := strings.NewReader("test\n")
		out := &bytes.Buffer{}
		game := &poker.GameSpy{}

		cli := poker.NewCLI(in, out, game)
		cli.PlayPoker()

		if game.StartCalled {
			t.Errorf("game shound not have started")
		}

		poker.AssertMessagesSentToUser(t, out, poker.PlayerPrompt+poker.BadPlayerInputErrMsg)
	})

}

func TestGame_Start(t *testing.T) {
	t.Run("schedules alerts on game start for 5 players", func(t *testing.T) {
		blindAlerter := &poker.SpyBlindAlerter{}
		game := poker.NewTexasHoldem(blindAlerter, dummyPlayerStore)

		game.Start(7, io.Discard)

		cases := []poker.ScheduledAlert{
			{At: 0 * time.Minute, Amount: 100},
			{At: 12 * time.Minute, Amount: 200},
			{At: 24 * time.Minute, Amount: 300},
			{At: 36 * time.Minute, Amount: 400},
		}

		checkSchedulingCases(cases, t, blindAlerter)
	})
}

func TestGame_Finish(t *testing.T) {
	t.Run("Finish game", func(t *testing.T) {
		store := &poker.StubPlayerStore{}
		game := poker.NewTexasHoldem(dummyBlindAlerter, store)
		winner := "Ruth"
		game.Finish(winner)
		poker.AssertPlayerWin(t, store, winner)
	})
}

func checkSchedulingCases(cases []poker.ScheduledAlert, t *testing.T, alerter *poker.SpyBlindAlerter) {
	t.Helper()
	for i, want := range cases {
		t.Run(fmt.Sprint(want), func(t *testing.T) {
			if len(alerter.Alerts) <= i {
				t.Fatalf("alert %d was not scheduled %v", i, alerter.Alerts)
			}

			got := alerter.Alerts[i]
			poker.AssertScheduledAt(t, got, want)
		})
	}
}
