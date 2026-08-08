package integers

import "testing"


func TestAdder(t *testing.T) {
	t.Run("base add", func(t *testing.T) {
		got:= Adder(12,15)
		want:= 27 
		assertMessage(t, got, want)
	})
}

func assertMessage(t testing.TB, got, want int) {
	t.Helper()
	if got!=want{
		t.Errorf("got %d want %d", got, want)
	}
}
