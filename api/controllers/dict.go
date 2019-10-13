package controllers

import (
	"github.com/linbaozhong/car-care/api/models"
)

//字典
type Dict struct {
	Base
}

// @Summary 字典-汽车外观检查项目
// @Description 汽车外观检查项目字典
// @Param appId query string true "appId"
// @Param tk query string true "token"
// @Success 200 {object} models.Result
// @Failure 400 {object} models.Result
// @router /waiguan [get]
func (d *Dict) Waiguan() {
	d.getDicts(models.NewDicts().Waiguan)
}

// @Summary 字典-汽车内饰检查项目
// @Description 汽车内饰检查项目字典
// @Param appId query string true "appId"
// @Param tk query string true "token"
// @Success 200 {object} models.Result
// @Failure 400 {object} models.Result
// @router /neishi [get]
func (d *Dict) Neishi() {
	d.getDicts(models.NewDicts().Neishi)
}

// @Summary 字典-汽车仪表检查项目
// @Description 汽车仪表检查项目字典
// @Param appId query string true "appId"
// @Param tk query string true "token"
// @Success 200 {object} models.Result
// @Failure 400 {object} models.Result
// @router /yibiao [get]
func (d *Dict) Yibiao() {
	d.getDicts(models.NewDicts().Yibiao)
}

// @Summary 字典-汽车其他检查项目
// @Description 汽车其他检查项目字典
// @Param appId query string true "appId"
// @Param tk query string true "token"
// @Success 200 {object} models.Result
// @Failure 400 {object} models.Result
// @router /qita [get]
func (d *Dict) Qita() {
	d.getDicts(models.NewDicts().Qita)
}

// @Summary 字典-汽车上线检查项目
// @Description
// @Param appId query string true "appId"
// @Param tk query string true "token"
// @Success 200 {object} models.Result
// @Failure 400 {object} models.Result
// @router /shangxian [get]
func (d *Dict) Shangxian() {
	d.getDicts(models.NewDicts().Shangxian)
}


func (d *Dict) getDicts(f func() ([]models.Dict, error)) {
	ds, err := f()
	d.Result.Data = ds
	if err != nil {
		d.Result.Message = err.Error()
	}
}
