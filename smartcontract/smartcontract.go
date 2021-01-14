package smartcontract

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"example.com/myfirstgo/entity/user"
	"encoding/json"
	// "log"
	
)

// SmartContract provides functions for managing archives
type SmartContract struct {
	contractapi.Contract
}

func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	users :=[]user.User{
		{Id:"1",Name:"user1",Email:"1@example.com",Description:"description1"},
		{Id:"2",Name:"user2",Email:"2@example.com",Description:"description2"},
		{Id:"3",Name:"user3",Email:"3@example.com",Description:"description3"},
	}
	
	for _,user := range users{
		
		userJSON, err := json.Marshal(user)
		if err != nil {
		  return err
		}
		err = ctx.GetStub().PutState(user.Id,userJSON)
		if err != nil {
			return err
		}
	}
	
	return nil
}
