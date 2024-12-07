package fabric

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/hash"
	"time"
)

type RepairProcess struct {
	Time    string `json:"time"`    //记录时间
	Content string `json:"content"` //维修内容
	Cost    string `json:"cost"`    // 维修费用
	Notes   string `json:"notes"`   // 备注
}

func AddRepairProcessRecord(repairId, processId, content, cost, notes string) error {
	clientConnection := newGrpcConnection()
	defer clientConnection.Close()

	identity := newIdentity()
	sign := newSign()

	gw, err := client.Connect(
		identity,
		client.WithSign(sign),
		client.WithHash(hash.SHA256),
		client.WithClientConnection(clientConnection),
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		return fmt.Errorf("failed to connect:%s", err.Error())
	}
	defer gw.Close()

	network := gw.GetNetwork(channel)
	contract := network.GetContract(repairChaincode)

	_, err = contract.SubmitTransaction("AddRepairProcessRecord", repairId, processId, content, cost, notes)
	if err != nil {
		return fmt.Errorf("failed to submit transaction:%s", err.Error())
	}
	return nil
}
func UpdateRepairProcessRecord(repairId, processId, content, cost, notes string) error {
	clientConnection := newGrpcConnection()
	defer clientConnection.Close()

	identity := newIdentity()
	sign := newSign()

	gw, err := client.Connect(
		identity,
		client.WithSign(sign),
		client.WithHash(hash.SHA256),
		client.WithClientConnection(clientConnection),
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		return fmt.Errorf("failed to connect:%s", err.Error())
	}
	defer gw.Close()

	network := gw.GetNetwork(channel)
	contract := network.GetContract(repairChaincode)

	_, err = contract.SubmitTransaction("UpdateRepairProcessRecord", repairId, processId, content, cost, notes)
	if err != nil {
		return fmt.Errorf("failed to submit transaction:%s", err.Error())
	}
	return nil
}

func GetRepairProcessRecordByRepairId(repairId string) ([]RepairProcess, error) {
	clientConnection := newGrpcConnection()
	defer clientConnection.Close()

	identity := newIdentity()
	sign := newSign()

	gw, err := client.Connect(
		identity,
		client.WithSign(sign),
		client.WithHash(hash.SHA256),
		client.WithClientConnection(clientConnection),
		client.WithEvaluateTimeout(5*time.Second),
		client.WithEndorseTimeout(15*time.Second),
		client.WithSubmitTimeout(5*time.Second),
		client.WithCommitStatusTimeout(1*time.Minute),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect:%s", err.Error())
	}
	defer gw.Close()

	network := gw.GetNetwork(channel)
	contract := network.GetContract(repairChaincode)

	result, err := contract.EvaluateTransaction("GetRepairProcessRecordByRepairId", repairId)
	if err != nil {
		return nil, fmt.Errorf("failed to submit transaction:%s", err.Error())
	}
	if result == nil {
		return nil, nil
	}

	var records []RepairProcess
	if err := json.Unmarshal(result, &records); err != nil {
		return nil, fmt.Errorf("failed to unmarshal result:%s", err.Error())
	}
	return records, nil
}
