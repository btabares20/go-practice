package iterations

import "strings"


func Repeat(letter string) string {
	count := 5
	var b strings.Builder;
	b.Grow(len(letter) * count)
	for range count {
		b.WriteString(letter)
	}
	return b.String()
}
