package post

import (
	// "example.com/myfirstgo/entity"
	// "encoding/json"
	// "log"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)
func GetPost(ctx contractapi.TransactionContextInterface, key string)(*Post,error){
	// postJSON, err := ctx.GetSub().GetState(key)
	// if err != nil {
	// 	log.Println("Get postJSON is Failed!!!")
	// 	return nil, err
	// }
	// if postJSON == nil {
	// 	return nil, entity.GetStateNotFoundError
	// }
	// var post Post
	// err=json.Unmarshal(postJSON,&post)
	// if err!= nil {
	// 	log.Println("Unmarshal is Failed!!!")
	// 	return nil, err
	// }
	// return &post,nil
	return nil,nil
}