package DataStructures

import "fmt"

func DataStructuresFour() {
	sampleKeyValues := map[string]int{
		"Key1": 1,
		"Key2": 2,
		"Key3": 3,
		"Key4": 4,
		"Key5": 5,
	}
	swappedKeyValues := make(map[int]string)
	for k, v := range sampleKeyValues {
		swappedKeyValues[v] = k
	}
	fmt.Println(swappedKeyValues)
}
