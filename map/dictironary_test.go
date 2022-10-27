package search

import (
	"testing"
)

func TestSearch(t *testing.T) {
	t.Run("Known word", func(t *testing.T) {
		key := "test"
		val := "this is just a test"
		dict := Dictionary{key: val}

		got, err := dict.Search(key)
		assertError(t, err, nil)
		assertString(t, got, val)
	})

	t.Run("Unknown word", func(t *testing.T) {
		dict := Dictionary{"test": "this is just a test"}

		_, err := dict.Search("testNotExist")

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("New word", func(t *testing.T) {
		dict := Dictionary{}
		key := "test"
		val := "this is just a test"
		dict.Add(key, val)

		assertValue(t, dict, key, val)

	})

	t.Run("Existing word", func(t *testing.T) {
		dict := Dictionary{}
		key := "test"
		val := "this is just a test"
		_ = dict.Add(key, val)
		err := dict.Add(key, val)

		assertError(t, err, ErrWordExists)
		assertValue(t, dict, key, val)
	})

}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		key := "test"
		val := "this is just a test"
		dict := Dictionary{key: val}
		newVal := "new value"

		err := dict.Update(key, newVal)

		assertError(t, err, nil)
		assertValue(t, dict, key, newVal)
	})

	t.Run("not existing word", func(t *testing.T) {
		key := "test"
		val := "this is just a test"
		dict := Dictionary{}

		err := dict.Update(key, val)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		key := "test"
		val := "this is just a test"
		dict := Dictionary{key: val}

		dict.Delete(key)

		_, err := dict.Search(key)
		if err != ErrNotFound {
			t.Errorf("Expected %q, to be deleted", key)
		}
	})
}

func assertString(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t *testing.T, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
	if got == nil {
		if want == nil {
			return
		}
		t.Fatal("expected to got an error")
	}
}

func assertValue(t *testing.T, d Dictionary, key, val string) {
	t.Helper()

	got, err := d.Search(key)

	if err != nil {
		t.Errorf("Not expected error")
	}

	if got != val {
		t.Errorf("got error %q want %q", got, val)
	}
}
