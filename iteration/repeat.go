package iteration

import "strings"

func Repeater(char string, length int) string{
	var result strings.Builder
	for range(length){
		result.WriteString(char) 
	}
	return result.String()
}
