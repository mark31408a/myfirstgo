package smartcontract

import (
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"example.com/myfirstgo/entity/user"
	// "example.com/myfirstgo/entity"
	"log"
	"encoding/json"
)

func (s *SmartContract)CreateUser(ctx contractapi.TransactionContextInterface, input user.CreateUserDto) (interface{}, error) {
	log.Println("########## CreateUser ##########")
	log.Println("input = ",input)

	newuser := user.User{
		Id:"4",
		Name:input.Name,
		Email:input.Email,
		Description:input.Description,
	}
	newuserJSON, err := json.Marshal(newuser)
    if err != nil {
      return nil,err
    }
	ctx.GetStub().PutState("4",newuserJSON)
	return nil,nil

}

func (s *SmartContract)GetAllUser(ctx contractapi.TransactionContextInterface) (interface{}, error) {
	log.Println("########## GetAllUser ##########")
	
	
	resultsIterator,err :=ctx.GetStub().GetStateByRange("", "")
	if err!=nil{
		return nil,err
	}
	var alluser []*user.User
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
		  return nil, err
		}

		var useri user.User
		err = json.Unmarshal(queryResponse.Value, &useri)
		if err != nil {
		  return nil, err
		}
		alluser = append(alluser, &useri)
	}
	alluserjson,err:=json.Marshal(alluser)
	if err!=nil{
		return nil,err
	}
	return string(alluserjson),nil

}
