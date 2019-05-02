package callisto

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func TestServe(t *testing.T) {
	go Serve()

	port := os.Getenv("PORT")

	resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%s/hello", port))

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

func TestStatic(t *testing.T) {
	go Serve()

	port := os.Getenv("PORT")

	resp, err := http.Get(fmt.Sprintf("http://127.0.0.1:%s", port))

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
