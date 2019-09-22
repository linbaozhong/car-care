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
	Level   int8   `json:"level" xorm:"'level'"`
	Pinyin  string `json:"pinyin" xorm:"'pinyin'"`
	First   string `json:"first" xorm:"'first'"`
}

func NewBrand() *Brand {
	return &Brand{}
}

func (m *Brand) TableName() string {
	return "brand"
}

func GetAllBrands() ([]*Brand, error) {
	ids, e := GetBrandIDs()
	if e != nil {
		return nil, e
	}
	ids_len := len(ids)
	if ids_len == 0 {
		return nil, errors.New("not found")
	}
	brands := make([]*Brand, 0, ids_len)
	for _, id := range ids {
		brand, _ := GetBrand(id)
		if brand == nil {
			continue
		}
		brands = append(brands, brand)
	}
	return brands, nil
}

func GetBrandIDs() ([]int, error) {
	var ids []int
	err := db.Table(&Brand{}).Cols("id").Find(&ids)
	return ids, err
}
func GetBrand(id int) (*Brand, error) {
	bean := NewBrand()
	_, err := db.ID(id).Get(bean)
	return bean, err
}
