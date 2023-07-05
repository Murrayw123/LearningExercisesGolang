package Testing

func TestingTwo(sampleInts [5]int) [5]int {
	for i := range sampleInts {
		sampleInts[i]++
	}
	return sampleInts
}
