package main

import (
	"errors"
	"fmt"
	"log"
	"util"
	"wrapper"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type PointsChaincode struct {
}

//初始化创建表
func (t *PointsChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	defer util.End(util.Begin("Init"))

	log.Println("Init method.........")
	return nil, util.CreateTable(stub)
}

//调用方法
func (t *PointsChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	defer util.End(util.Begin("Invoke"))

	if function == "CreditPoints" { // 授信积分
		log.Println("CreditPoints,args = " + args[0])
		return wrapper.CreditPoints(stub, args)
	} else if function == "ConsumePoints" { // 消费积分
		log.Println("ConsumePoints,args = " + args[0])
		return wrapper.ConsumePoints(stub, args)
	} else if function == "AccpetPoints" { // 承兑积分
		log.Println("AccpetPoints,args = " + args[0])
		return wrapper.AccpetPoints(stub, args)
	} else if function == "InitData" {
		log.Println("InitData,args = " + args[0])
		return wrapper.InitData(stub, args)
	}

	return nil, errors.New("调用invoke失败")
}

//查询
func (t *PointsChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	defer util.End(util.Begin("Query"))
	//通过表名查条数
	if function == "QueryTableLines" {
		log.Println("QueryTableLines,args = " + args[0])
		fmt.Println("QueryTableLines,args = " + args[0])
		return wrapper.QueryTableLines(stub, args)
	} else if function == "QueryPointsByKey" { //通过主键查询次条记录
		log.Println("QueryPointsByKey,args = " + args[0])
		fmt.Println("QueryPointsByKey,args = " + args[0])
		return wrapper.QueryPointsByKey(stub, args)
	}

	return nil, errors.New("调用query失败")
}

func main() {

	err := shim.Start(new(PointsChaincode))
	if err != nil {
		fmt.Printf("Error starting chaincode: %s", err)
	}

}
