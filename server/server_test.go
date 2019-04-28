package callisto

import "testing"

func TestHello(t *testing.T) {
	res := Hello()
	if res != "Hello" {
		t.Error("result should be \"Hello\".")
	}
}
