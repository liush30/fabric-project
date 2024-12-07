package fabric

import "testing"

func TestRegisterPile(t *testing.T) {
	err := RegisterPile("1", "2-HASH")
	if err != nil {
		t.Error(err)
	}
}
func TestQueryPileHistory(t *testing.T) {
	pile, err := QueryPileHistory("1")
	if err != nil {
		t.Error(err)
	}

	t.Log(pile)
}
