package DataStructures

import "fmt"

func contains(slice []int, val int) bool {
	for _, v := range slice {
		if v == val {
			return true
		}
	}
	return false
}

func DataStructuresTwo() {
	sliceOne := []int{1, 3, 2, 3, 3, 4, 5, 5, 6, 7, 8, 9}
	deduped := make([]int, 0)
	for _, v := range sliceOne {
		if !contains(deduped, v) {
			deduped = append(deduped, v)
		}
	}
	fmt.Println(deduped)
}
