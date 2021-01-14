package test

import (
	// "example.com/myfirstgo/entity/post"
	"example.com/myfirstgo/smartcontract"
	// "encoding/json"
	//  "fmt"
	// "io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/gogo/protobuf/proto"
	// "github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-chaincode-go/shimtest"
	"github.com/hyperledger/fabric-contract-api-go/contractapi"
	"github.com/hyperledger/fabric-protos-go/msp"
)

const certWithAttrs = `-----BEGIN CERTIFICATE-----
MIIC2TCCAn+gAwIBAgIUQ0IZAeWJyRqPFpcFshvpVbY1RzMwCgYIKoZIzj0EAwIw
ZjELMAkGA1UEBhMCVVMxFzAVBgNVBAgTDk5vcnRoIENhcm9saW5hMRQwEgYDVQQK
EwtIeXBlcmxlZGdlcjEPMA0GA1UECxMGY2xpZW50MRcwFQYDVQQDEw5yY2Etb3Jn
MS1hZG1pbjAeFw0xODExMTMxNzQ4MDBaFw0xOTExMTMxNzUzMDBaMG8xCzAJBgNV
BAYTAlVTMRcwFQYDVQQIEw5Ob3J0aCBDYXJvbGluYTEUMBIGA1UEChMLSHlwZXJs
ZWRnZXIxHDANBgNVBAsTBmNsaWVudDALBgNVBAsTBG9yZzExEzARBgNVBAMTCmFk
bWluLW9yZzEwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAR196Xv7te+C5gkz7Ui
h8t2gl8QjjSs6iOLFTk18IEH5vLh+DovGT9q3ylvZpExtOap5zFkCva9GnChxP05
4A0eo4IBADCB/TAOBgNVHQ8BAf8EBAMCB4AwDAYDVR0TAQH/BAIwADAdBgNVHQ4E
FgQUXf9wjawRl/KosmHcVnYB4ay8IqswHwYDVR0jBBgwFoAUwqQ3h+jBjt2e2wC1
f1amDdCHY7QwFwYDVR0RBBAwDoIMZjExN2MxODEyYzM3MIGDBggqAwQFBgcIAQR3
eyJhdHRycyI6eyJhYmFjLmluaXQiOiJ0cnVlIiwiYWRtaW4iOiJ0cnVlIiwiaGYu
QWZmaWxpYXRpb24iOiJvcmcxIiwiaGYuRW5yb2xsbWVudElEIjoiYWRtaW4tb3Jn
MSIsImhmLlR5cGUiOiJjbGllbnQifX0wCgYIKoZIzj0EAwIDSAAwRQIhAN1v/XK0
WmZf5u9X9FG5uGxwcJ9d5K/eFAC7KahSbs65AiB/GzS2u1cYznXzTDWoBm9oflxY
w8Ou1Sh9IjeXj/SDAA==
-----END CERTIFICATE-----
`

var Stub *shimtest.MockStub
var Scc *contractapi.ContractChaincode
var MSPID = "MSP99"


func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}

func setup() {
	log.SetOutput(os.Stdout)//ioutil.Discard)
}

func NewStub() {
	Scc, err := contractapi.NewChaincode(new(smartcontract.SmartContract))
	if err != nil {
		log.Println("NewChaincode failed:", err)
		os.Exit(0)
	}
	
	Stub = shimtest.NewMockStub("main", Scc)

	SetMspID(MSPID)
}

func SetMspID(mspid string) {
	if Stub == nil {
		log.Println("Stub not initial")
		os.Exit(0)
	}

	sid := &msp.SerializedIdentity{Mspid: mspid, IdBytes: []byte(certWithAttrs)}
	data, err := proto.Marshal(sid)
	if err != nil {
		log.Println("Marshal false")
		os.Exit(0)
	}

	Stub.Creator = data
}

func InitLedger(){
	Stub.MockInvoke("uuid", [][]byte{[]byte("InitLedger")})
}