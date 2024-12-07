package main

import (
	"github.com/hyperledger/fabric-contract-api-go/v2/contractapi"
	"log"
)

func main() {
	contract, err := contractapi.NewChaincode(&PileContract{})
	if err != nil {
		log.Panicf("error create pile contracr:%s", err.Error())
	}
	if err := contract.Start(); err != nil {
		log.Panicf("failed start chaincode:%s", err.Error())
	}

}
