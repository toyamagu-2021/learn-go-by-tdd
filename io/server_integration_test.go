package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	db, cleanDatbase := createTempFile(t, `[]`)
	defer cleanDatbase()
	store, err := NewFileSystemPlayerStore(db)

	assertNoError(t, err)

	server := NewPlayerServer(store)
	player := "Pepper"

	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

	t.Run("Get score", func(t *testing.T) {
		res := httptest.NewRecorder()
		server.ServeHTTP(res, newGetScoreRequest(player))
		assertStatus(t, res.Code, http.StatusOK)

		assertResponseBody(t, res.Body.String(), "3")
	})

	t.Run("get league", func(t *testing.T) {
		res := httptest.NewRecorder()
		server.ServeHTTP(res, newLeagueRequest())
		assertStatus(t, res.Code, http.StatusOK)

		got := getLeagueFromResponse(t, res.Body)
		want := []Player{
			{"Pepper", 3},
		}

		assertLeague(t, got, want)
	})

}
