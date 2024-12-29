package model

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGenerateJson(t *testing.T) {
	data, err := json.Marshal(&Station{
		StationID:     "",
		RepairmanID:   "",
		StationName:   "",
		Location:      "",
		City:          "",
		District:      "",
		ContactNumber: "",
		ManagerName:   "",
		OpeningHours:  "",
		Status:        0,
		Description:   "",
		LoginPwd:      "",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(string(data))
}
