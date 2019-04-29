package callisto

import "testing"

func TestHello(t *testing.T) {
	res := Hello("Aspirin")

	if res != "Hello, Aspirin" {
		t.Errorf("Should be \"Hello, Aspirin\", but got: %s", res)
	}
}
