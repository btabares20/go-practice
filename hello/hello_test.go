package hello

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	t.Run("Test Hello base", func(t *testing.T){
		got := Hello("", "english")
		want := "Hello, World"
		assert(t, want, got)
	});
	t.Run("Test Hello with name", func(t *testing.T){
		got := Hello("Bryan", "english")
		want := "Hello, Bryan"
		assert(t, want, got)
	});
	t.Run("Test Hello spanish base", func(t *testing.T){
		got := Hello("", "spanish")
		want := "Hola, World"
		assert(t, want, got)
	});
	t.Run("Test Hello with name, spanish", func(t *testing.T){
		got := Hello("Bryan", "spanish")
		want := "Hola, Bryan"
		assert(t, want, got)
	});
	t.Run("Test Hello french base", func(t *testing.T){
		got := Hello("", "french")
		want := "Bonjour, World"
		assert(t, want, got)
	});
	t.Run("Test Hello with name", func(t *testing.T){
		got := Hello("Bryan", "french")
		want := "Bonjour, Bryan"
		assert(t, want, got)
	});
}

func assert(t testing.TB, want, got string){
	t.Helper()
	if want != got {
		t.Errorf("want %q, but got %q", want, got)
	}
}

func ExampleHello(){
	got := Hello("", "english")
	fmt.Println(got)
	// Output: Hello, World
}
