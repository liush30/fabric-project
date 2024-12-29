package model

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGenerateJson(t *testing.T) {
	data, err := json.Marshal(&Repairman{
		RepairmanID: "1",
		UserName:    "1",
		Password:    "1",
		Name:        "1",
		ContactInfo: "1",
		Status:      0,
		Description: "1",
		UserType:    0,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(data))
}
