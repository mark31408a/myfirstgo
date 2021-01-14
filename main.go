package main

import (
	"log"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"example.com/myfirstgo/smartcontract"
)
func main(){
	assetChaincode, err := contractapi.NewChaincode(&smartcontract.SmartContract{})
	if err != nil {
		log.Panicf("Error creating asset-transfer-basic chaincode: %v", err)
	}

	if err := assetChaincode.Start(); err != nil {
		log.Panicf("Error starting asset-transfer-basic chaincode: %v", err)
	}
	return
}