package captain

import (
	"io/ioutil"
	"testing"
	"time"
)

func TestGetAccounts(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	client := newClientFromEnv(t)
	accounts, err := client.GetAccounts(withTimeout(time.Second * 5))
	if err != nil {
		t.Fatal(err)
	}
	accountUUID := mustGetenv(t, "CAPTAIN_ACCOUNT_ID")
	found := false
	for _, account := range accounts {
		if *account.UUID == accountUUID {
			found = true
		}
	}
	if !found {
		t.Errorf("expected provided account in list")
	}
}

func TestParseAccount(t *testing.T) {
	data, err := ioutil.ReadFile("testdata/account.json")
	if err != nil {
		t.Fatal(err)
	}
	account := &Account{
		UUID:         String("100000ben1es"),
		FriendlyName: String("Bennies Pizza"),
	}
	testExactJSON(t, account, data)
}
