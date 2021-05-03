package integers

import "testing"

func TestAdder(t *testing.T){
	assertCorrectMessage := func(t testing.TB, got, want int) {
        t.Helper()
        if got != want {
            t.Errorf("got %q want %q", got, want)
        }
    }

	t.Run("Simple Add 2 + 2", func(t *testing.T)  {
		got := Add(2, 2)
		want := 4
		assertCorrectMessage(t, got, want)
	})
}