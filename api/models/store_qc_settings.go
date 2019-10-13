package models

import (
	"errors"
	"github.com/linbaozhong/car-care/api/utils"
	"sync"
	"time"
)

var (
	cache_store_qc_settings      = utils.NewCache("store_qc_settings", 30*time.Minute)
	locker_store_qc_settings     = new(sync.Mutex)
)
//StoreQCSetting
type StoreQCSetting struct {
	ID      int    `json:"id" xorm:"pk autoincr 'id'"`
	StoreID int    `json:"store_id" xorm:"'store_id'"`
	Type    int8   `json:"type" xorm:"'type'"`
	Items   string `json:"items" xorm:"'items'"`
}

func NewStoreQCSetting() *StoreQCSetting {
	return &StoreQCSetting{}
}

func (*StoreQCSetting) TableName() string {
	return "store_qc_settings"
}

func GetStoreQCSettings(store_id int64) ([]StoreQCSetting, error) {
	if val, ok := cache_store.Get(store_id); ok {
		if _val, _ok := val.([]StoreQCSetting); _ok {
			return _val, nil
		}
		return nil, errors.New("cache类型错误")
	}

	locker_store.Lock()
	defer locker_store.Unlock()

	if val, ok := cache_store.Get(store_id); ok {
		if _val, _ok := val.([]StoreQCSetting); _ok {
			return _val, nil
		}
		return nil, errors.New("cache类型错误")
	}
	beans := make([]StoreQCSetting,0)
	err := db.Where("`store_id`=?",store_id).Find(&beans)
	cache_store.Set(store_id, beans)
	return beans, err
}
