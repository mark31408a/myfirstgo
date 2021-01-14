package user

import (
	"example.com/myfirstgo/entity"
	"encoding/json"
	"log"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)
func GetUser(ctx contractapi.TransactionContextInterface, id string)(*User,error){
	userJSON, err := ctx.GetStub().GetState(id)
	if err != nil {
		log.Println("Get postJSON is Failed!!!")
		return nil, err
	}
	if userJSON == nil {
		return nil, entity.GetStateNotFoundError
	}
	var user User
	err=json.Unmarshal(userJSON,&user)
	if err!= nil {
		log.Println("Unmarshal is Failed!!!")
		return nil, err
	}
	return &user,nil
	return nil,nil
}