package wrapper

import (
	"account"
	"errors"
	"log"
	"points"
	"util"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type CreditPointsTransData struct {
	Account                 *account.Account
	PointsTransaction       *points.PointsTransaction
	PointsTransactionDetail *points.PointsTransactionDetail
	AuditObj                *util.AuditObject
}

type ConsumePointsTransData struct {
	AccountList                 []*account.Account
	PointsTransaction           *points.PointsTransaction
	PointsTransactionDetailList []*points.PointsTransactionDetail
	AuditObj                    *util.AuditObject
}
type AccpetPointsTransData struct {
	AccountList                 []*account.Account
	PointsTransaction           *points.PointsTransaction
	PointsTransactionDetailList []*points.PointsTransactionDetail
	AuditObj                    *util.AuditObject
}

//注册
func SignUp(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	return nil, nil
}

//登录
func SignIn(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	return nil, nil
}

//授信积分
func CreditPoints(stub shim.ChaincodeStubInterface, args []string) error {
	// 解析传入数据
	creditObject := new(CreditPointsTransData)
	err := util.ParseJsonAndDecode(creditObject, args)
	if err != nil {
		log.Println("Error occurred when parsing json")
		return errors.New("Error occurred when parsing json.")
	}

	//账户信息表更新
	if creditObject.PointsTransaction.OperFlag == "1" {
		err := account.UpdateAccount(stub, creditObject.Account)
		if err != nil {
			log.Println("Error occurred when performing UpdateAccount")
			return errors.New("Error occurred when performing UpdateAccount.")
		}

	} else {
		log.Println("Flag is not 1,that's why we do nothing for table account")
	}

	// 积分交易表增加记录
	if creditObject.Account.OperFlag == "0" {
		err := points.InsertPointsTransation(stub, creditObject.PointsTransaction)
		if err != nil {
			log.Println("Error occurred when performing InsertPointsTransation")
			return errors.New("Error occurred when performing InsertPointsTransation.")
		}

	} else {
		log.Println("Flag is not 0,that's why we do nothing for table points_transation")
	}

	// 积分交易明细表增加记录
	if creditObject.Account.OperFlag == "0" {
		err := points.InsertPointsTransationDetail(stub, creditObject.PointsTransactionDetail)
		if err != nil {
			log.Println("Error occurred when performing InsertPointsTransationDetail")
			return errors.New("Error occurred when performing InsertPointsTransationDetail.")
		}

	} else {
		log.Println("Flag is not 0,that's why we do nothing for table points_transation_detail")
	}

	log.Println("credit points is ok!!")

	return nil
}

//消费积分
func ConsumePoints(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	// 解析传入数据
	data := new(ConsumePointsTransData)
	err := util.ParseJsonAndDecode(data, args)
	if err != nil {
		log.Println("Error occurred when parsing json")
		return nil, errors.New("Error occurred when parsing json.")
	}
	for i := 0; i < len(data.AccountList); i++ {
		//如果标识符为0就对账户表新增否则修改
		if data.AccountList[i].OperFlag == "0" {
			account.InsertAccount(stub, *data.AccountList[i])
		} else {
			account.UpdateAccount(stub, data.AccountList[i])
		}
	}
	//如果标识符为0就对积分交易表新增否则报错返回
	if data.PointsTransaction.OperFlag == "0" {
		points.InsertPointsTransation(stub, data.PointsTransaction)
	} else {
		return nil, errors.New("输入标识符有误")
	}
	for i := 0; i < len(data.PointsTransactionDetailList); i++ {
		if data.PointsTransactionDetailList[i].OperFlag == "0" {
			points.InsertPointsTransationDetail(stub, data.PointsTransactionDetailList[i])
		} else {
			points.UpdatePointsTransationDetail(stub, data.PointsTransactionDetailList[i])
		}
	}

	return nil, nil
}

//承兑积分
func AccpetPoints(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	// 解析传入数据
	data := new(AccpetPointsTransData)
	err := util.ParseJsonAndDecode(data, args)
	if err != nil {
		log.Println("Error occurred when parsing json")
		return nil, errors.New("Error occurred when parsing json.")
	}
	for i := 0; i < len(data.AccountList); i++ {
		//如果标识符为0就对账户表新增否则修改
		if data.AccountList[i].OperFlag == "0" {
			account.InsertAccount(stub, *data.AccountList[i])
		} else {
			account.UpdateAccount(stub, data.AccountList[i])
		}
	}
	//如果标识符为0就对积分交易表新增否则报错返回
	if data.PointsTransaction.OperFlag == "0" {
		points.InsertPointsTransation(stub, data.PointsTransaction)
	} else {
		return nil, errors.New("输入标识符有误")
	}
	for i := 0; i < len(data.PointsTransactionDetailList); i++ {
		if data.PointsTransactionDetailList[i].OperFlag == "0" {
			points.InsertPointsTransationDetail(stub, data.PointsTransactionDetailList[i])
		} else {
			points.UpdatePointsTransationDetail(stub, data.PointsTransactionDetailList[i])
		}
	}
	return nil, nil
}
