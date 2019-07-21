package captain

import (
	"testing"
	"time"
)

func TestAuth(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	client := newClientFromEnv(t)
	testAuthResponse, err := client.TestAuth(withTimeout(time.Second * 5))
	if err != nil {
		t.Fatal(err)
	}
	logJSON(t, testAuthResponse)
	if testAuthResponse.Message != "Successfully Authorised Using Developer Key and Integration Key." {
		t.Errorf("unexpected response message")
	}
}
