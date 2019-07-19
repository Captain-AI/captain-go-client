package captain

import (
	"context"
	"os"
	"testing"
	"time"
)

func TestLiveRequests(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	accountID := os.Getenv("CAPTAIN_ACCOUNT_ID")
	if accountID == "" {
		t.Fatal("missing environment variable: CAPTAIN_ACCOUNT_ID")
	}
	integrationKey := os.Getenv("CAPTAIN_INTEGRATION_KEY")
	if integrationKey == "" {
		t.Fatal("missing environment variable: CAPTAIN_INTEGRATION_KEY")
	}
	developerKey := os.Getenv("CAPTAIN_DEVELOPER_KEY")
	if developerKey == "" {
		t.Fatal("missing environment variable: CAPTAIN_DEVELOPER_KEY")
	}
	client := NewClient()
	client.IntegrationKey = integrationKey
	client.DeveloperKey = developerKey
	testAuthResponse, err := client.TestAuth(withTimeout(time.Second * 5))
	if err != nil {
		t.Fatal(err)
	}
	if testAuthResponse.Message != "Successfully Authorised Using Developer Key and Integration Key." {
		t.Errorf("unexpected response message")
	}
}

func withTimeout(timeout time.Duration) context.Context {
	ctxWithTimeout, _ := context.WithTimeout(context.Background(), timeout)
	return ctxWithTimeout
}
