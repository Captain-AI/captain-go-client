package captain

import (
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/go-test/deep"
)

func newClientFromEnv(t *testing.T) *Client {
	t.Helper()
	client := NewClient()
	client.SetIntegrationKey(mustGetenv(t, "CAPTAIN_INTEGRATION_KEY"))
	client.SetDeveloperKey(mustGetenv(t, "CAPTAIN_DEVELOPER_KEY"))
	return client
}

func mustGetenv(t *testing.T, name string) string {
	t.Helper()
	value := os.Getenv(name)
	if value == "" {
		t.Fatalf("missing environment variable: %q", name)
	}
	return value
}

func withTimeout(timeout time.Duration) context.Context {
	ctxWithTimeout, _ := context.WithTimeout(context.Background(), timeout)
	return ctxWithTimeout
}

func logJSON(t *testing.T, v interface{}) {
	t.Helper()
	data, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(data))
}

func testExactJSON(t *testing.T, v interface{}, data []byte) {
	t.Helper()
	var want interface{}
	err := json.Unmarshal(data, &want)
	if err != nil {
		t.Fatal(err)
	}
	buf, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	var have interface{}
	err = json.Unmarshal(buf, &have)
	if err != nil {
		t.Fatal(err)
	}
	if diff := deep.Equal(want, have); diff != nil {
		t.Fatal(diff)
	}
}
