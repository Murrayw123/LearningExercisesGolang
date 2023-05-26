package Testing

import "testing"

func TestTestingTwo(t *testing.T) {
	res := TestingTwo("Hello")
	if res != "olleH" {
		t.Errorf("Expected Hello, got %s", res)
	}
}
