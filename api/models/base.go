package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	jsoniter "github.com/json-iterator/go"
	"sync"
)

var (
	db         *xorm.Engine
	resultPool = sync.Pool{
		New: func() interface{} {
			m := Result{
				Status: 200,
			}
			return &m
		},
	}
	json = jsoniter.ConfigCompatibleWithStandardLibrary
)

func init() {
	var err error
	db, err = xorm.NewEngine("mysql", "root:123456@/cars?charset=utf8mb4")
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	db.ShowSQL(true)
}

const (
	QC_TYPE_WAIGUAN   = 1 //汽车外观检查项目
	QC_TYPE_NEISHI    = 2 //汽车内饰检查项目
	QC_TYPE_YIBIAO    = 3 //汽车仪表检查项目
	QC_TYPE_QITA      = 4 //汽车其它检查项目
	QC_TYPE_SHANGXIAN = 5 //汽车上线检查项目

	ROLE_BOSS    = 10 //老板
	ROLE_MANAGER = 20 //店长
	ROLE_CLERK   = 30 //店员

	SECURITY_CODE = "2019-10-01"
)
