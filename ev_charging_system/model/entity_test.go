package model

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGenerateJson(t *testing.T) {
	data, err := json.Marshal(&Pile{
		PileID:      "",
		StationID:   "",
		PileCode:    "",
		PileName:    "",
		Description: "",
		Location:    "",
		Status:      0,
		Type:        0,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(data))
}
