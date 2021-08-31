package attack_check_server

import "testing"

func TestWAFAttackCheck(t *testing.T) {
	InitTestDB()
	err := waf{}.WAFAttackCheck()
	if err != nil {
		t.Fatal(err)
	}
}
