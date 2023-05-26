package DataStructures

import (
	"fmt"
	"strings"
)

func DataStructuresThree() {
	words := "Word 1, Word 2, Word 1, Word 3, Word 1, Word 2, Word 4, Word 2"
	wordsSlice := strings.Split(words, ", ")
	wordsMap := make(map[string]int)
	for _, v := range wordsSlice {
		wordsMap[v]++
	}
	fmt.Println(wordsMap)
}
