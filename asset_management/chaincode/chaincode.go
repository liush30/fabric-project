package chaincode

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
)

type AssetContract struct {
	contractapi.Contract
}

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

const FormatId = "auction_%s"

// InitAsset 初始化资产
func (a *AssetContract) InitAsset(ctx contractapi.TransactionContextInterface, assetId, hash, owner, notes string) error {
	//判断记录是否存在
	exist, err := ctx.GetStub().GetState(assetId)
	if err != nil {
		return fmt.Errorf("the record already exists")
	}
	if exist != nil {
		return fmt.Errorf("the record already exists")
	}
	asset := Asset{
		AssetHash: hash,
		Notes:     notes,
		Owner:     owner,
	}
	assetBytes, err := json.Marshal(asset)
	if err != nil {
		return fmt.Errorf("failed to marshal record: %v", err)
	}

	return ctx.GetStub().PutState(assetId, assetBytes)
}

// UploadAssessmentResult 上传评估结果
func (a *AssetContract) UploadAssessmentResult(ctx contractapi.TransactionContextInterface, assetId, assessor, result, note string) error {
	//判断记录是否存在
	exist, err := ctx.GetStub().GetState(assetId)
	if err != nil {
		return fmt.Errorf("the record already exists")
	}
	if exist == nil {
		return fmt.Errorf("the record not exists")
	}
	//获取时间
	time, err := ctx.GetStub().GetTxTimestamp()
	if err != nil {
		return fmt.Errorf("failed to get current time: %v", err)
	}
	formatStr := time.AsTime().Format("2006-01-02 15:04:05")
	assessment := AssessmentResult{
		Assessor:         assessor,
		AssessmentTime:   formatStr,
		AssessmentResult: result,
		AssessmentNote:   note,
	}
	var asset Asset
	if err := json.Unmarshal(exist, &asset); err != nil {
		return fmt.Errorf("failed to unmarshal record: %v", err)
	}
	asset.AssessmentResult = assessment

	assetBytes, err := json.Marshal(asset)
	if err != nil {
		return fmt.Errorf("failed to marshal record: %v", err)
	}
	return ctx.GetStub().PutState(assetId, assetBytes)
}

// UpdateOwner 更新资产所有者
func (a *AssetContract) UpdateOwner(ctx contractapi.TransactionContextInterface, assetId, owner, notes string) error {
	exist, err := ctx.GetStub().GetState(assetId)
	if err != nil {
		return fmt.Errorf("the record already exists")
	}
	if exist == nil {
		return fmt.Errorf("the record not exists")
	}
	var asset Asset
	if err := json.Unmarshal(exist, &asset); err != nil {
		return fmt.Errorf("failed to unmarshal record: %v", err)
	}
	asset.Owner = owner
	asset.Notes = notes
	assetBytes, err := json.Marshal(asset)
	if err != nil {
		return fmt.Errorf("failed to marshal record: %v", err)
	}
	return ctx.GetStub().PutState(assetId, assetBytes)
}

// CreateAuction 创建拍卖任务
func (a *AssetContract) CreateAuction(ctx contractapi.TransactionContextInterface, assetId, auctionId, assetHash, result, notes string) error {
	//判断记录是否存在
	exist, err := ctx.GetStub().GetState(assetId)
	if err != nil {
		return fmt.Errorf("the record already exists")
	}
	if exist == nil {
		return fmt.Errorf("the record not exists")
	}
	auction := Auction{
		AuctionId: auctionId,
		AssetHash: assetHash,
		Notes:     notes,
		Result:    result,
	}
	auctionBytes, err := json.Marshal(auction)
	if err != nil {
		return fmt.Errorf("failed to marshal record: %v", err)
	}
	return ctx.GetStub().PutState(fmt.Sprintf(FormatId, assetId), auctionBytes)
}

// QueryAuctionHistory 查询拍卖任务历史记录
func (a *AssetContract) QueryAuctionHistory(ctx contractapi.TransactionContextInterface, assetId string) ([]Auction, error) {
	resultsIterator, err := ctx.GetStub().GetHistoryForKey(fmt.Sprintf(FormatId, assetId))
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	var auctions []Auction
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var auction Auction
		if err := json.Unmarshal(response.Value, &auction); err != nil {
			return nil, err
		}
		auctions = append(auctions, auction)
	}
	return auctions, nil
}

// QueryAssetHistory 查询资产历史记录
func (a *AssetContract) QueryAssetHistory(ctx contractapi.TransactionContextInterface, assetId string) ([]Asset, error) {
	resultsIterator, err := ctx.GetStub().GetHistoryForKey(assetId)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()
	var assets []Asset
	for resultsIterator.HasNext() {
		response, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		var asset Asset
		if err := json.Unmarshal(response.Value, &asset); err != nil {
			return nil, err
		}
		assets = append(assets, asset)
	}
	return assets, nil
}

// 查询资产信息
func (a *AssetContract) QueryAsset(ctx contractapi.TransactionContextInterface, assetId string) (*Asset, error) {
	exist, err := ctx.GetStub().GetState(assetId)
	if err != nil {
		return nil, fmt.Errorf("the record already exists")
	}
	if exist == nil {
		return nil, fmt.Errorf("the record not exists")
	}
	var asset Asset
	if err := json.Unmarshal(exist, &asset); err != nil {
		return nil, fmt.Errorf("failed to unmarshal record: %v", err)
	}
	return &asset, nil
}
