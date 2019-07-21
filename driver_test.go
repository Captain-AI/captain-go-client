package captain

import (
	"testing"
	"time"
)

func TestGetDrivers(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	client := newClientFromEnv(t)
	accountUUID := mustGetenv(t, "CAPTAIN_ACCOUNT_ID")

	// TODO: API is broken right now ...
	t.Skip()

	drivers, err := client.GetDrivers(withTimeout(time.Second*5), accountUUID)
	if err != nil {
		t.Fatal(err)
	}
	logJSON(t, drivers)
	t.Fail()
}
