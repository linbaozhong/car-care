package models

import (
	"errors"
	"sync"
)

var (
	brand_ids_locker = new(sync.Mutex)
)
//Brand 品牌
type Brand struct {
	ID      int    `json:"id" xorm:"pk autoincr 'id'"`
	Name    string `json:"name" xorm:"'name'"`
	LogoUrl string `json:"logo_url" xorm:"'logo_url'"`
	Maker   string `json:"maker" xorm:"'maker'"`
	Sort    int8   `json:"sort" xorm:"'sort'"`
	Pinyin  string `json:"pinyin" xorm:"'pinyin'"`
	First   string `json:"first" xorm:"'first'"`
}

func NewBrand() *Brand {
	return &Brand{}
}

func (m *Brand) TableName() string {
	return "brand"
}

//GetAllBrands 读取全部品牌
func GetAllBrands() ([]Brand, error) {
	ids, e := GetAllBrandIDs()
	if e != nil {
		return nil, e
	}
	return getBrands(ids)
}

//GetAllBrandIDS 读取全部品牌ids
func GetAllBrandIDs() ([]int, error) {
	var ids []int
	err := db.Table(&Brand{}).OrderBy("first").Cols("id").Find(&ids)
	return ids, err
}
func GetUsageBrands() ([]Brand, error) {
	ids, e := GetUsageBrandIDS()
	if e != nil {
		return nil, e
	}
	return getBrands(ids)
}

//GetUsageBrandIDS 读取常用品牌
func GetUsageBrandIDS() ([]int, error) {
	var ids []int
	err := db.Table(&Brand{}).Where("`sort`>0").OrderBy("sort").Cols("id").Find(&ids)
	return ids, err
}

func getBrands(ids []int) ([]Brand, error) {
	ids_len := len(ids)
	if ids_len == 0 {
		return nil, errors.New("not found")
	}
	brands := make([]Brand, 0, ids_len)
	for _, id := range ids {
		brand, _ := GetBrand(id)
		if brand == nil {
			continue
		}
		brands = append(brands, *brand)
	}
	return brands, nil
}
func GetBrand(id int) (*Brand, error) {
	bean := NewBrand()
	_, err := db.ID(id).Get(bean)
	return bean, err
}
