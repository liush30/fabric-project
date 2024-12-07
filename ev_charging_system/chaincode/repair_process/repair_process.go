package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)

type RepairProcessContract struct {
	contractapi.Contract
}
type RepairProcess struct {
	Time    string `json:"time"`    //记录时间
	Content string `json:"content"` //维修内容
	Cost    string `json:"cost"`    // 维修费用
	Notes   string `json:"notes"`   // 备注
}

func (r *RepairProcessContract) AddRepairProcessRecord(ctx contractapi.TransactionContextInterface, repairId, processId, content, cost, notes string) error {
	//设置复合键
	key, err := ctx.GetStub().CreateCompositeKey(repairId, []string{repairId, processId})
	if err != nil {
		return fmt.Errorf("failed to create composite key: %v", err)
	}
	//判断key记录是否存在
	exist, err := ctx.GetStub().GetState(key)
	if err != nil {
		return fmt.Errorf("the record already exists")
	}
	if exist != nil {
		return fmt.Errorf("the record already exists")
	}
	// 获取当前时间
	time, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return fmt.Errorf("failed to get current time: %v", err)
	}
	formatStr := time.AsTime().Format("2006-01-02 15:04:05")

	record := RepairProcess{
		Time:    formatStr,
		Content: content,
		Cost:    cost,
		Notes:   notes,
	}
	recordBytes, err := json.Marshal(record)
	if err != nil {
		return fmt.Errorf("failed to marshal record: %v", err)
	}

	return ctx.GetStub().PutState(key, recordBytes)
}

// UpdateRepairProcessRecord 更新维修记录
func (r *RepairProcessContract) UpdateRepairProcessRecord(ctx contractapi.TransactionContextInterface, repairId, processId, content, cost, notes string) error {
	//设置复合键
	key, err := ctx.GetStub().CreateCompositeKey(repairId, []string{repairId, processId})
	if err != nil {
		return fmt.Errorf("failed to create composite key: %v", err)
	}
	//判断记录是否存在
	exist, err := ctx.GetStub().GetState(key)
	if err != nil {
		return fmt.Errorf("the record already exists")
	}
	if exist == nil {
		return fmt.Errorf("the record does not exist")
	}
	//获取当前时间
	time, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return fmt.Errorf("failed to get current time: %v", err)
	}
	formatStr := time.AsTime().Format("2006-01-02 15:04:05")
	record := RepairProcess{
		Time:    formatStr,
		Content: content,
		Cost:    cost,
		Notes:   notes,
	}
	recordBytes, err := json.Marshal(record)
	if err != nil {
		return fmt.Errorf("failed to marshal record: %v", err)
	}
	return ctx.GetStub().PutState(key, recordBytes)
}

// GetRepairProcessRecordByRepairId 指定repairId获取维修记录
func (r *RepairProcessContract) GetRepairProcessRecordByRepairId(ctx contractapi.TransactionContextInterface, repairId string) ([]RepairProcess, error) {
	resultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(repairId, []string{repairId})
	if err != nil {
		return nil, fmt.Errorf("failed to get history: %v", err)
	}
	defer resultsIterator.Close()
	var records []RepairProcess
	for resultsIterator.HasNext() {
		record, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to get next record: %v", err)
		}
		var repairProcess RepairProcess
		err = json.Unmarshal(record.Value, &repairProcess)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal record: %v", err)
		}
		records = append(records, repairProcess)
	}
	return records, nil
}
