package Testing

import (
	"fmt"
	"strings"
)

func TestingOne(sampleString string) string {
	newString := make([]string, len(sampleString))
	for i, v := range sampleString {
		fmt.Println(i, string(v))
		newString[len(sampleString)-i-1] = string(v)
	}
	return strings.Join(newString, "")
}
