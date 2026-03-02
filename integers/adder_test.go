package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T){
	t.Run("Test normal addition", func(t *testing.T){
		sum := Adder(2,2)
		expected := 4
		assertEqualInt(t, sum, expected)
	})
}

func assertEqualInt(t testing.TB, want, got int){
	if want!=got {
		t.Errorf("want %q, got %q", want, got)
	}
}

func ExampleAdder(){
	sum := Adder(1, 5)
	fmt.Println(sum)
	// Output: 6
}
