package main

import "testing"

func TestSearch(t *testing.T) {
	dic := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dic.Search("test")
		want := "this is just a test"

		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dic.Search("unknown")
		want := ErrNotFound

		assertError(t, got, want)
	})
}

func TestAdd(t *testing.T) {
	word := "test"
	definition := "this is just a test"

	t.Run("new word", func(t *testing.T) {
		dic := Dictionary{}
		dic.Add(word, definition)

		want := definition
		got, err := dic.Search(word)

		assertError(t, err, nil)
		assertString(t, got, want)
	})

	t.Run("existing word", func(t *testing.T) {
		dic := Dictionary{word: definition}
		err := dic.Add(word, "new meaning")

		want := definition
		got, _ := dic.Search(word)

		assertError(t, err, ErrWordExists)
		assertString(t, want, got)
	})
}

func TestUpdate(t *testing.T) {
	word := "test"
	definition := "this is just a test"

	t.Run("existing word", func(t *testing.T) {
		dic := Dictionary{word: definition}
		newDefinition := "new defenition"

		dic.Update(word, newDefinition)
		got, _ := dic.Search(word)
		want := newDefinition

		assertString(t, got, want)
	})

	t.Run("new word", func(t *testing.T) {
		dic := Dictionary{}

		err := dic.Update(word, definition)
		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	word := "test"
	dic := Dictionary{word: "test def"}

	dic.Delete(word)

	_, err := dic.Search(word)
	assertError(t, err, ErrNotFound)
}

func assertString(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got != want {
		t.Errorf("got error %q, want error %q", got, want)
	}
}
