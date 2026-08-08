package iterations

import "testing"


func TestRepeat(t *testing.T) {
	t.Run("base repeat", func(t *testing.T) {
		got:= Repeat("a")
		want:= "aaaaa"
		assertMessage(t, got, want)
	})
}
func BenchmarkRepeat(b *testing.B) {
	for b.Loop() {
		Repeat("a")
	}
}
func assertMessage(t testing.TB, got, want string) {
	t.Helper()
	if got!=want{
		t.Errorf("got %q want %q", got, want)
	}
}
