package iteration

import "testing"

func TestRepeater(t *testing.T){
	result := Repeater("a", 10)
	expected := "aaaaaaaaaa"
	assertEqualString(t, expected, result)
}
func BenchmarkRepeater(b *testing.B){
	for b.Loop(){
		Repeater("a", 10)
	}
}

func assertEqualString(t testing.TB, want, got string){
	t.Helper()
	if want != got {
		t.Errorf("want %q, got %q", want, got)
	}
}
