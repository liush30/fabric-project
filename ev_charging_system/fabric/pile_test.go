package fabric

import "testing"

func TestRegisterPile(t *testing.T) {
	err := RegisterPile("1", "2-HASH")
	if err != nil {
		t.Error(err)
	}
}
func TestQueryPileHistory(t *testing.T) {
	pile, err := QueryGunHistory("406693dfa73948d0a22ab22ded711be0", "452770781d454cba8d3beffdabbce69e")
	if err != nil {
		t.Error(err)
	}

	t.Log(pile)
}

func TestDeletePile(t *testing.T) {
	err := DeletePile("3d673a9637fb463b95c0a7807aea79db")
	if err != nil {
		t.Error(err)
	}
}
