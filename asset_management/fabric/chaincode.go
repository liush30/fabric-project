package fabric

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-gateway/pkg/client"
	"github.com/hyperledger/fabric-gateway/pkg/hash"
	"time"
)

type Asset struct {
	AssetHash        string           `json:"hash"`              //资产信息hash
	AssessmentResult AssessmentResult `json:"assessment_result"` //资产评估结果
	Owner            string           `json:"owner"`             //当前所有人
	Notes            string           `json:"notes"`             //备注
}

// 拍卖任务记录
type Auction struct {
	AuctionId string `json:"auction_id"` //拍卖任务id
	AssetHash string `json:"hash"`       //拍卖任务ID
	Result    string `json:"result"`     //拍卖结果
	Notes     string `json:"notes"`      //备注
}

// AssessmentResult  评估结果
type AssessmentResult struct {
	//评估人
	Assessor string `json:"assessor"`
	//评估时间
	AssessmentTime string `json:"assessment_time"`
	//评估结果
	AssessmentResult string `json:"result"`
	//评估说明
	AssessmentNote string `json:"note"`
}

func InitAsset(assetId, assetHash, owner, notes string) error {
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
	contract := network.GetContract(pileChaincode)
	_, err = contract.SubmitTransaction("InitAsset", assetId, assetHash, owner, notes)
	if err != nil {
		return fmt.Errorf("failed to submit transaction:%s", err.Error())
	}
	return nil
}

func UploadAssessmentResult(assetId, assessor, result, note string) error {
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
	contract := network.GetContract(pileChaincode)

	_, err = contract.SubmitTransaction("UploadAssessmentResult", assetId, assessor, result, note)
	if err != nil {
		return fmt.Errorf("failed to submit transaction:%s", err.Error())
	}
	return nil
}

func UpdateOwner(assetId, owner, notes string) error {
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
	contract := network.GetContract(pileChaincode)

	_, err = contract.SubmitTransaction("UpdateOwner", assetId, owner, notes)
	if err != nil {
		return fmt.Errorf("failed to submit transaction:%s", err.Error())
	}
	return nil
}

func CreateAuction(assetId, auctionId, assetHash, result, notes string) error {
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
	contract := network.GetContract(pileChaincode)

	_, err = contract.SubmitTransaction("CreateAuction", assetId, auctionId, assetHash, result, notes)
	if err != nil {
		return fmt.Errorf("failed to submit transaction:%s", err.Error())
	}
	return nil
}

func QueryAuctionHistory(assetId string) ([]Auction, error) {
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
	contract := network.GetContract(pileChaincode)

	result, err := contract.EvaluateTransaction("QueryAuctionHistory", assetId)
	if err != nil {
		return nil, fmt.Errorf("failed to submit transaction:%s", err.Error())
	}
	if result == nil {
		return nil, nil
	}

	var auctions []Auction
	err = json.Unmarshal(result, &auctions)
	if err != nil {
		return nil, fmt.Errorf("failed to submit transaction:%s", err.Error())
	}
	return auctions, nil
}

func QueryAssetHistory(assetId string) ([]Asset, error) {
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
	contract := network.GetContract(pileChaincode)

	result, err := contract.EvaluateTransaction("QueryAssetHistory", assetId)
	if err != nil {
		return nil, fmt.Errorf("failed to submit transaction:%s", err.Error())
	}
	if result == nil {
		return nil, nil
	}

	var assets []Asset
	err = json.Unmarshal(result, &assets)
	if err != nil {
		return nil, fmt.Errorf("failed to submit transaction:%s", err.Error())
	}
	return assets, nil

}

func QueryAsset(assetId string) (Asset, error) {
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
		return Asset{}, fmt.Errorf("failed to connect:%s", err.Error())
	}
	defer gw.Close()

	network := gw.GetNetwork(channel)
	contract := network.GetContract(pileChaincode)

	result, err := contract.EvaluateTransaction("QueryAsset", assetId)
	if err != nil {
		return Asset{}, fmt.Errorf("failed to submit transaction:%s", err.Error())
	}

	var asset Asset
	err = json.Unmarshal(result, &asset)
	if err != nil {
		return Asset{}, fmt.Errorf("failed to submit transaction:%s", err.Error())
	}
	return asset, nil
}
