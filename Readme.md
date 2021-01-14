# my first chaincode in go
## chain code package
#### check go vertion and requierments
```shell
cat go.mod
```
#### test chain code (with log)
```shell
go test -v test/*
```
#### test chain code (without log)
```shell
go test test/*
```
#### run main (not neccesary)
```shell
go run main.go
```
>Cause the code is designed to run on a peer of ledger, you might have some error while running main. Don't worry about it.
#### download go packages locally
```go mod vendor```
>the go packages will be installed inside a vendor folder, thus, we can deploy to the ledger.
---
## deploy chaincode to test network
>download fabric sample from https://github.com/hyperledger/fabric-samples and go through the tutorial in https://hyperledger-fabric.readthedocs.io/en/release-2.2/tutorials.html

>>>The content below is a brief copy of https://hyperledger-fabric.readthedocs.io/en/release-2.2/deploy_chaincode.html
#### Bring up the test network
go inside the fabric sample repository
```shell
cd test-network 
./network.sh up createChannel -ca
``` 

#### Package chaincode
```shell
export PATH=${PWD}/../bin:$PATH
export FABRIC_CFG_PATH=$PWD/../config/
peer lifecycle chaincode package basic.tar.gz --path PATH_OF_YOUR_CHAINCODE --lang golang --label basic_1.0
```
replace {PATH_OF_YOUR_CHAINCODE} with the relative path from "fabric-sample/test-network" to "myfirstgo/". for example if the two repository are in the same diretory it will be "../../myfirstgo".

the whole package will pack into basic.tar.gz inside "fabric-sample/test-network" 

we are using basic as the name of chaincode, take it easy to try another name.

#### Install the chaincode package


##### set to org1
```shell
export CORE_PEER_TLS_ENABLED=true
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_ADDRESS=localhost:7051
```
```shell
peer lifecycle chaincode install basic.tar.gz
```
##### set to org2
```shell
export CORE_PEER_LOCALMSPID="Org2MSP"
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org2.example.com/users/Admin@org2.example.com/msp
export CORE_PEER_ADDRESS=localhost:9051
```
```shell
peer lifecycle chaincode install basic.tar.gz
```
#### Approve chaincode
```shell
peer lifecycle chaincode queryinstalled
```
checking the ID of chaincode
```shell
Installed chaincodes on peer:
Package ID: basic_1.0:69de748301770f6ef64b42aa6bb6cb291df20aa39542c3ef94008615704007f3, Label: basic_1.0
```
!!!IMPORTANT!!! 
copy and paste the id you see in you terminal to the following command
```shell
export CC_PACKAGE_ID=basic_1.0:69de748301770f6ef64b42aa6bb6cb291df20aa39542c3ef94008615704007f3
```
```shell
peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --channelID mychannel --name basic --version 1.0 --package-id $CC_PACKAGE_ID --sequence 1 --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
```
also approve with another org(in this case it's org1)
```shell
export CORE_PEER_LOCALMSPID="Org1MSP"
export CORE_PEER_MSPCONFIGPATH=${PWD}/organizations/peerOrganizations/org1.example.com/users/Admin@org1.example.com/msp
export CORE_PEER_TLS_ROOTCERT_FILE=${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt
export CORE_PEER_ADDRESS=localhost:7051
peer lifecycle chaincode approveformyorg -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --channelID mychannel --name basic --version 1.0 --package-id $CC_PACKAGE_ID --sequence 1 --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
```
#### Commit chainode
```shell
peer lifecycle chaincode commit -o localhost:7050 --ordererTLSHostnameOverride orderer.example.com --channelID mychannel --name basic --version 1.0 --sequence 1 --tls --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem --peerAddresses localhost:7051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org1.example.com/peers/peer0.org1.example.com/tls/ca.crt --peerAddresses localhost:9051 --tlsRootCertFiles ${PWD}/organizations/peerOrganizations/org2.example.com/peers/peer0.org2.example.com/tls/ca.crt
```
#### Check 
```shell
peer lifecycle chaincode querycommitted --channelID mychannel --name basic --cafile ${PWD}/organizations/ordererOrganizations/example.com/orderers/orderer.example.com/msp/tlscacerts/tlsca.example.com-cert.pem
```
if everyihing goes right you will see something like
```shell
Committed chaincode definition for chaincode 'basic' on channel 'mychannel':
Version: 1.0, Sequence: 1, Endorsement Plugin: escc, Validation Plugin: vscc, Approvals: [Org1MSP: true, Org2MSP: true]
```
now we can use the deployed chaincode through other application such as a backend server().