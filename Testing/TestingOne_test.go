package Testing

import "testing"

func TestTestingOne(t *testing.T) {
	res := TestingOne("Hello")
	if res != "olleH" {
		t.Errorf("Expected Hello, got %s", res)
	}
}
