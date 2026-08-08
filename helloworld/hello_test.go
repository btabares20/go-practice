package helloworld 

import "testing"

func TestHello(t *testing.T){
	t.Run("named hello", func(t *testing.T) {
		got:= Hello("Bryan","")
		want:= "Hello, Bryan"
		assertMessage(t, got, want)
	})
	t.Run("default hello", func(t *testing.T) {
		got:= Hello("","")
		want:= "Hello, World"
		assertMessage(t, got, want)
	})
	t.Run("spanish hello", func(t *testing.T) {
		got:= Hello("","spanish")
		want:= "Hola, World"
		assertMessage(t, got, want)
	})
	t.Run("spanish hello with name", func(t *testing.T) {
		got:= Hello("Bryan","spanish")
		want:= "Hola, Bryan"
		assertMessage(t, got, want)
	})
	t.Run("wrong hello", func(t *testing.T) {
		got:= Hello("","french")
		want:= "Hello, World"
		assertMessage(t, got, want)
	})
	t.Run("wrong hello with name", func(t *testing.T) {
		got:= Hello("Bryan","french")
		want:= "Hello, Bryan"
		assertMessage(t, got, want)
	})
}
func assertMessage(t testing.TB, got, want string) {
	t.Helper()
	if got!=want{
		t.Errorf("got %q want %q", got, want)
	}
}
