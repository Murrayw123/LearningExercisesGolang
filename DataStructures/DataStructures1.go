package DataStructures

import (
	"fmt"
	"strconv"
	"strings"
)

func DataStructuresOne() {
	strArray := "1, 2, 3, 4, 5, 6, 7"

	// convert strArray to an array of ints
	strArray = strings.ReplaceAll(strArray, " ", "")
	slice := strings.Split(strArray, ",")

	total := 0

	for i, str := range slice {
		num, err := strconv.Atoi(str)
		if err != nil {
			fmt.Printf("Error converting string to int at index %d: %v\n", i, err)
			return
		}
		total += num
	}

	fmt.Println("Total: ", total)
}
