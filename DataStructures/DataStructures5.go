package DataStructures

import (
	"strings"
)

func DataStructuresFive() {
	sampleWords := "Alphabet, Banana, Aardvark, Carrot, Apple, Zebra, Carrot, Banana"
	sampleWordsSlice := strings.Split(sampleWords, ", ")
	orderedSlice := make([]string, len(sampleWords))

	// sort the ordered slice alphabetically
	for _, v := range sampleWordsSlice {
		for i, v2 := range orderedSlice {
			if v2 == "" {
				orderedSlice[i] = v
				break
			} else if v < v2 {
				orderedSlice = append(orderedSlice[:i], append([]string{v}, orderedSlice[i:]...)...)
				break
			}
		}
	}

}
