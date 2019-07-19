package captain

import (
	"encoding/json"
	"reflect"
	"testing"
)

func TestParseAccounts(t *testing.T) {
	response := `[{"uuid":"100000ben1es","friendly_name":"Bennies Pizza"},{"uuid":"8013a9468e78","friendly_name":"Bennies London"}]`
	haveAccounts := []*Account{}
	err := json.Unmarshal([]byte(response), &haveAccounts)
	if err != nil {
		t.Fatal(err)
	}
	wantAccounts := []*Account{
		&Account{UUID: "100000ben1es", FriendlyName: "Bennies Pizza"},
		&Account{UUID: "8013a9468e78", FriendlyName: "Bennies London"},
	}
	if !reflect.DeepEqual(haveAccounts, wantAccounts) {
		t.Errorf("account: incorrect json struct tags")
	}
}
