package test

import (
	"testing"
	"log"
    "example.com/myfirstgo/entity/user"
    "encoding/json"
    
)


func TestGetUser(t *testing.T) {
    NewStub()
    InitLedger()
    res:= Stub.MockInvoke("uuid", [][]byte{[]byte("GetAllUser")})
    // if res==nil{
    //     log.Println("no user exist")
	// 	t.FailNow()
    // }
    users := [3]user.User {
        {
            Id:"1",
            Name:"user1",
            Email:"1@example.com",
            Description:"description1",
        },
        {
            Id:"2",
            Name:"user2",
            Email:"2@example.com",
            Description:"description2",
        },
        {
            Id:"3",
            Name:"user3",
            Email:"3@example.com",
            Description:"description3",
        },
    }
    want, err := json.Marshal(users)
    
    if err != nil {
        t.Errorf("Error: %s", err)
        return;
    }
    if got := string(res.Payload); got != string(want) {

        t.Errorf("user = %s, want %s", got, string(want))
    }
}

func TestCreateUser(t *testing.T){
    NewStub()
    InitLedger()
    usernew := user.User {
        Id:"4",
        Name:"usernew",
        Email:"new@example.com",
        Description:"descriptionnew",
    }
    createusernewInput:= user.CreateUserDto{
        Name:usernew.Name,
        Email:usernew.Email,
        Description:usernew.Description,
    }
    createusernewJSON ,err := json.Marshal(createusernewInput)
    if err!=nil{
        log.Println("Failed to Marshal")
		t.FailNow()
    }
    Stub.MockInvoke("uuid", [][]byte{[]byte("CreateUser"),createusernewJSON})

    usergetJSON,err := Stub.GetState(usernew.Id)
    if err!=nil{
        log.Println("Failed to GetState after Create")
		t.FailNow()
    }
    if usergetJSON==nil{
        log.Println("user not exist")
		t.FailNow()
    }
    var userget user.User
    err = json.Unmarshal(usergetJSON, &userget)
    if err!=nil{
        log.Println("Failed to unmarshal",err)
		t.FailNow()
    }
    if userget != usernew {
        t.Errorf("userget = %q, want %q", userget, usernew)
    }
    
}