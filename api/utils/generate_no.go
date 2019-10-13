package utils

import (
	"github.com/astaxie/beego/logs"
	"github.com/bwmarrin/snowflake"
	"sync"
)

var (
	generateNodes       = make(map[string]*snowflake.Node)
	lockerGenerateNodes sync.Mutex
)

func getGenerateNode(key string) *snowflake.Node {
	node := generateNodes[key]
	if node != nil {
		return node
	}
	
	lockerGenerateNodes.Lock()
	defer lockerGenerateNodes.Unlock()
	
	node = generateNodes[key]
	if node != nil {
		return node
	}
	
	node, err := snowflake.NewNode(int64(len(generateNodes) + 1))
	if err != nil {
		logs.Error("编号生成器错误：%s", err.Error())
	}
	generateNodes[key] = node
	
	return node
}

//GetOrderNo 读取一个新的订单编号
func GetOrderNo() int64 {
	node := getGenerateNode("order")
	return node.Generate().Int64()
}

//GetPaymentNo 生成一个新的编号
func GetPaymentNo() int64 {
	node := getGenerateNode("paylog")
	return node.Generate().Int64()
}
