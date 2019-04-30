package callisto

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestServe(t *testing.T) {
	go Serve()

	resp, err := http.Get("http://127.0.0.1:3000/hello")

	if err != nil {
		t.Error(err)
	} else if resp.StatusCode != http.StatusOK {
		t.Errorf("Status Code not ok: %d", resp.StatusCode)
	}

	defer resp.Body.Close()

	res, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		t.Errorf("Read Error: %v", err)
	}

	t.Logf("Result is: %s", res)
}
