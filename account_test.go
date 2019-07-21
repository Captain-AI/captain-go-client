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
	logJSON(t, accounts)
	if len(accounts) == 0 {
		t.Errorf("expected accounts")
	}
}

/*
func TestGetAccount(t *testing.T) {
	if testing.Short() {
		t.Skip()
	}
	client := newClientFromEnv(t)
	accountUUID := mustGetenv(t, "CAPTAIN_ACCOUNT_ID")
	account, err := client.GetAccount(withTimeout(time.Second*5), accountUUID)
	if err != nil {
		t.Fatal(err)
	}
	if account.UUID == nil {
		t.Errorf("expected account UUID")
	}
}
*/

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
