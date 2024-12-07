package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)

type PileContract struct {
	contractapi.Contract
}

type Pile struct {
	PileHash string `json:"pileHash"` // 充电桩hash
	Time     string `json:"time"`     // 记录时间
	Option   string `json:"option"`   // 操作
}
type Gun struct {
	GunHash string `json:"gunHash"` // 充电枪hash
	Time    string `json:"time"`    // 记录时间
	Option  string `json:"option"`  // 操作
}

const (
	OptionRegister = "register"
	OptionUpdate   = "update"
	OptionDelete   = "delete"
)

// RegisterPile 注册充电桩
func (p *PileContract) RegisterPile(ctx contractapi.TransactionContextInterface, id, pileHash string) error {
	//判断记录是否存在
	exist, err := ctx.GetStub().GetState(id)
	if err != nil {
		return fmt.Errorf("the record already exists")
	}
	if exist != nil {
		return fmt.Errorf("the record already exists")
	}
	//获取当前时间
	time, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return fmt.Errorf("failed to get current time: %v", err)
	}
	formatStr := time.AsTime().Format("2006-01-02 15:04:05")
	pile := Pile{
		PileHash: pileHash,
		Time:     formatStr,
		Option:   OptionRegister,
	}
	pileBytes, err := json.Marshal(pile)
	if err != nil {
		return fmt.Errorf("failed to marshal record: %v", err)
	}
	return ctx.GetStub().PutState(id, pileBytes)
}

// UpdatePile 更新充电桩
func (p *PileContract) UpdatePile(ctx contractapi.TransactionContextInterface, id, pileHash string) error {
	//判断记录是否存在
	exist, err := ctx.GetStub().GetState(id)
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
	pile := Pile{
		PileHash: pileHash,
		Time:     formatStr,
		Option:   OptionUpdate,
	}
	pileBytes, err := json.Marshal(pile)
	if err != nil {
		return fmt.Errorf("failed to marshal record: %v", err)
	}
	return ctx.GetStub().PutState(id, pileBytes)
}

// DeletePile 删除充电桩
func (p *PileContract) DeletePile(ctx contractapi.TransactionContextInterface, id, pileHash string) error {
	// 判断记录是否存在
	exist, err := ctx.GetStub().GetState(id)
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
	pile := Pile{
		PileHash: pileHash,
		Time:     formatStr,
		Option:   OptionDelete,
	}
	pileBytes, err := json.Marshal(pile)
	if err != nil {
		return fmt.Errorf("failed to marshal record: %v", err)
	}
	return ctx.GetStub().PutState(id, pileBytes)
}

// RegisterGun 注册充电枪
func (p *PileContract) RegisterGun(ctx contractapi.TransactionContextInterface, pileId, gunId, gunHash string) error {
	//判断pile记录是否存在
	pileExist, err := ctx.GetStub().GetState(pileId)
	if err != nil {
		return fmt.Errorf("the record already exists")
	}
	if pileExist == nil {
		return fmt.Errorf("the record not exists")
	}
	//判断key记录是否存在
	key, err := ctx.GetStub().CreateCompositeKey(pileId, []string{pileId, gunId})
	if err != nil {
		return fmt.Errorf("failed to create composite key: %v", err)
	}
	exist, err := ctx.GetStub().GetState(key)
	if err != nil {
		return fmt.Errorf("the record already exists")
	}
	if exist != nil {
		return fmt.Errorf("the record already exists")
	}
	//获取当前时间
	time, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return fmt.Errorf("failed to get current time: %v", err)
	}
	formatStr := time.AsTime().Format("2006-01-02 15:04:05")
	gun := Gun{
		GunHash: gunHash,
		Time:    formatStr,
		Option:  OptionRegister,
	}
	gunBytes, err := json.Marshal(gun)
	if err != nil {
		return fmt.Errorf("failed to marshal record: %v", err)
	}
	return ctx.GetStub().PutState(key, gunBytes)
}

// UpdateGun 更新充电枪
func (p *PileContract) UpdateGun(ctx contractapi.TransactionContextInterface, pileId, gunId, gunHash string) error {
	//判断pile记录是否存在
	pileExist, err := ctx.GetStub().GetState(pileId)
	if err != nil {
		return fmt.Errorf("the record already exists")
	}
	if pileExist == nil {
		return fmt.Errorf("the record not exists")
	}
	//判断key记录是否存在
	key, err := ctx.GetStub().CreateCompositeKey(pileId, []string{pileId, gunId})
	if err != nil {
		return fmt.Errorf("failed to create composite key: %v", err)
	}
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
	gun := Gun{
		GunHash: gunHash,
		Time:    formatStr,
		Option:  OptionUpdate,
	}
	gunBytes, err := json.Marshal(gun)
	if err != nil {
		return fmt.Errorf("failed to marshal record: %v", err)
	}
	return ctx.GetStub().PutState(key, gunBytes)
}

// DeleteGun 删除充电枪
func (p *PileContract) DeleteGun(ctx contractapi.TransactionContextInterface, pileId, gunId string) error {
	//判断pile记录是否存在
	pileExist, err := ctx.GetStub().GetState(pileId)
	if err != nil {
		return fmt.Errorf("the record already exists")
	}
	if pileExist == nil {
		return fmt.Errorf("the record not exists")
	}
	//判断key记录是否存在
	key, err := ctx.GetStub().CreateCompositeKey(pileId, []string{pileId, gunId})
	if err != nil {
		return fmt.Errorf("failed to create composite key: %v", err)
	}
	exist, err := ctx.GetStub().GetState(key)
	if err != nil {
		return fmt.Errorf("the record already exists")
	}
	if exist == nil {
		return fmt.Errorf("the record does not exist")
	}
	return ctx.GetStub().DelState(key)
}

// QueryPileHistory 查询充电桩历史变更记录
func (p *PileContract) QueryPileHistory(ctx contractapi.TransactionContextInterface, id string) ([]Pile, error) {
	resultsIterator, err := ctx.GetStub().GetHistoryForKey(id)
	if err != nil {
		return nil, fmt.Errorf("failed to get history: %v", err)
	}
	defer resultsIterator.Close()
	var records []Pile
	for resultsIterator.HasNext() {
		record, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to get next record: %v", err)
		}
		var pile Pile
		err = json.Unmarshal(record.Value, &pile)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal record: %v", err)
		}
		records = append(records, pile)
	}
	return records, nil
}

// QueryGunByPile 查看指定的所有充电枪信息
func (p *PileContract) QueryGunByPile(ctx contractapi.TransactionContextInterface, pileId string) ([]Gun, error) {
	resultsIterator, err := ctx.GetStub().GetStateByPartialCompositeKey(pileId, []string{pileId})
	if err != nil {
		return nil, fmt.Errorf("failed to get history: %v", err)
	}
	defer resultsIterator.Close()
	var records []Gun
	for resultsIterator.HasNext() {
		record, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to get next record: %v", err)
		}
		var gun Gun
		err = json.Unmarshal(record.Value, &gun)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal record: %v", err)
		}
		records = append(records, gun)
	}
	return records, nil
}

// QueryGunHistory 查看指定充电桩的充电枪历史变更记录
func (p *PileContract) QueryGunHistory(ctx contractapi.TransactionContextInterface, pileId, gunId string) ([]Gun, error) {
	key, err := ctx.GetStub().CreateCompositeKey(pileId, []string{pileId, gunId})
	if err != nil {
		return nil, fmt.Errorf("failed to create composite key: %v", err)
	}
	resultsIterator, err := ctx.GetStub().GetHistoryForKey(key)
	if err != nil {
		return nil, fmt.Errorf("failed to get history: %v", err)
	}
	defer resultsIterator.Close()
	var records []Gun
	for resultsIterator.HasNext() {
		record, err := resultsIterator.Next()
		if err != nil {
			return nil, fmt.Errorf("failed to get next record: %v", err)
		}
		var gun Gun
		err = json.Unmarshal(record.Value, &gun)
		if err != nil {
			return nil, fmt.Errorf("failed to unmarshal record: %v", err)
		}
		records = append(records, gun)
	}
	return records, nil
}
