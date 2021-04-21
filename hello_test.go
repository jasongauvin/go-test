package hello

import "testing"

func Test_main(t *testing.T) {
	got := hello("Jason")
	want := "Hello, Jason"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
