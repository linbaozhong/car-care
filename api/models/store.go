package models

import (
	"errors"
	"github.com/linbaozhong/car-care/api/utils"
	"sync"
	"time"
)

var (
	cache_store      = utils.NewCache("store", 30*time.Minute)
	locker_store     = new(sync.Mutex)
	cache_store_ids  = utils.NewCache("store:ids", 30*time.Minute)
	locker_store_ids = new(sync.Mutex)
)

//Store
type Store struct {
	ID         int     `json:"id" xorm:"pk autoincr 'id'"`
	Name       string  `json:"name" xorm:"'name'"`
	StartTime  string  `json:"start_time" xorm:"'start_time'"`
	EndTime    string  `json:"end_time" xorm:"'end_time'"`
	Address    string  `json:"address" xorm"'address'"`
	Longitude  float32 `json:"longitude" xorm"'longitude'"`
	Latitude   float32 `json:"latitude" xorm:"'latitude'"`
	Tel        string  `json:"tel" xorm:"'tel'"`
	Descr      string  `json:"descr" xorm:"'descr'"`
	CreateTime int64   `json:"create_time" xorm:"'create_time'"`
}

func (*Store) TableName() string {
	return "store"
}

func NewStore() *Store {
	return &Store{}
}

func GetStore(id int64) (*Store, error) {
	if val, ok := cache_store.Get(id); ok {
		if _val, _ok := val.(*Store); _ok {
			return _val, nil
		}
		return nil, errors.New("cache类型错误")
	}

	locker_store.Lock()
	defer locker_store.Unlock()

	if val, ok := cache_store.Get(id); ok {
		if _val, _ok := val.(*Store); _ok {
			return _val, nil
		}
		return nil, errors.New("cache类型错误")
	}
	bean := NewStore()
	has, err := db.Id(id).Get(bean);
	if !has {
		cache_store.Set(id, nil)
		return nil, err
	}
	cache_store.Set(id, bean)
	return bean, nil
}
