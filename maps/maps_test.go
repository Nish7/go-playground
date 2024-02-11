package maps

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("known word", func(t *testing.T) {

		got, _ := dictionary.Search("test")
		want := "this is a test"

		assertString(t, got, want)

	})

	t.Run("unknown word", func(t *testing.T) {

		_, err := dictionary.Search("tss")

		assertError(t, err, ErrNotFound)
	})

}

func TestAdd(t *testing.T) {
	dictionary := Dictionary{"test": "val"}

	t.Run("new word", func(t *testing.T) {
		dictionary.Add("new", "val")
		got, err := dictionary.Search("new")

		want := "val"

		if err != nil {
			t.Fatal("should find the added word:", err)
		}

		assertString(t, got, want)
	})

	t.Run("existing word", func(t *testing.T) {
		dictionary.Add("test", "val")
		got, err := dictionary.Search("new")

		want := "val"

		if err != nil {
			t.Fatal("should find the added word:", err)
		}

		assertString(t, got, want)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("update exisiting word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is a test"}

		newDef := "new val"
		dictionary.Update("test", newDef)

		got, err := dictionary.Search("test")

		if err != nil {
			t.Fatal("Unexpected Error")
		}

		if got != newDef {
			t.Errorf("got %s wanted %s, given %v", got, newDef, dictionary)
		}

	})

}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{"test": "val"}

	err := dictionary.Delete("test")

	if err != nil {
		t.Fatalf(err.Error())
	}

	_, err = dictionary.Search("test")

	if err != ErrNotFound {
		t.Errorf("Expected %q to be deleted", "test")
	}

}

func assertString(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
