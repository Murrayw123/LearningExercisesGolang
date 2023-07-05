package Testing

import (
	"fmt"
	"testing"
)

func TestTestingTwo(t *testing.T) {
	sampleInts := [5]int{1, 2, 3, 4, 5}

	fmt.Println("Test 1", sampleInts)
	res := TestingTwo(sampleInts)
	fmt.Println("Test 2", sampleInts)

	for i, v := range res {
		if v != sampleInts[i]+1 {
			t.Errorf("Expected %d, got %d", sampleInts[i]+1, v)
		}
	}
}
