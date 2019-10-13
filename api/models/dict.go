package models

import (
	"github.com/linbaozhong/car-care/api/utils"
	"sync"
	"time"
)

type Dict struct {
	Key  int    `json:"value"`
	Name string `json:"label"`
}

type Dicts struct {
	ID   int `json:"id"`
	Type int `json:"type"`
	Dict `xorm:"extends"`
}

func NewDicts() *Dicts {
	return &Dicts{}
}

var (
	createDictLocker = new(sync.Mutex)
	dictCaches       map[int]*utils.Cache
	dictListCaches   = utils.NewCache("dict_list_set", 1000*time.Hour)
	dictLocker       map[int]sync.Locker
	dictListLocker   map[int]sync.Locker
	dictCacheKey     = []string{
		"1_dict_waiguan",
		"2_dict_neishi",
		"3_dict_yibiao",
		"4_dict_qita",
		"5_dict_shangxian",
	}

	dictBrandsLocker = new(sync.Mutex)
	dictModelsLocker = new(sync.Mutex)
)

func init() {
	dictCaches = make(map[int]*utils.Cache)
	n := len(dictCacheKey)
	for i := 0; i < n; i++ {
		dictCaches[i+1] = utils.NewCache(dictCacheKey[i], 1000*time.Hour)
	}

	dictLocker = make(map[int]sync.Locker)
	for i := 0; i < n; i++ {
		dictLocker[i+1] = new(sync.Mutex)
	}

	dictListLocker = make(map[int]sync.Locker)
	for i := 0; i < n; i++ {
		dictListLocker[i+1] = new(sync.Mutex)
	}
}

func (d *Dicts) TableName() string {
	return "dict"
}

/* 汽车外观检查项目 */
// 读取 汽车外观检查项目 名称
func (d *Dicts) GetWaiguan(key int) string {
	return d.get(QC_TYPE_WAIGUAN, key)
}

// 读取 汽车外观检查项目 字典
func (d *Dicts) Waiguan() ([]Dict, error) {
	return d.query(QC_TYPE_WAIGUAN)
}

/* 汽车内饰检查项目 */
// 读取 汽车内饰检查项目 名称
func (d *Dicts) GetNeishi(key int) string {
	return d.get(QC_TYPE_NEISHI, key)
}

// 读取 汽车内饰检查项目 字典
func (d *Dicts) Neishi() ([]Dict, error) {
	return d.query(QC_TYPE_NEISHI)
}

/* 汽车仪表检查项目 */
// 读取 汽车仪表检查项目 名称
func (d *Dicts) GetYibiao(key int) string {
	return d.get(QC_TYPE_YIBIAO, key)
}

// 读取 汽车仪表检查项目 字典
func (d *Dicts) Yibiao() ([]Dict, error) {
	return d.query(QC_TYPE_YIBIAO)
}

//汽车其他检查项目
// 读取 汽车其他检查项目 名称
func (d *Dicts) GetQita(key int) string {
	return d.get(QC_TYPE_QITA, key)
}

//读取 汽车其他检查项目 字典
func (d *Dicts) Qita() ([]Dict, error) {
	return d.query(QC_TYPE_QITA)
}

//汽车上线检查项目
// 读取 汽车上线检查项目 名称
func (d *Dicts) GetShangxian(key int) string {
	return d.get(QC_TYPE_SHANGXIAN, key)
}

//读取 汽车上线检查项目 字典
func (d *Dicts) Shangxian() ([]Dict, error) {
	return d.query(QC_TYPE_SHANGXIAN)
}

// 创建新字典
func (d *Dicts) createDict(t int, name string) bool {
	createDictLocker.Lock()
	defer createDictLocker.Unlock()
	//如果找到，返回成功
	has, _ := db.Where("`type`=? and `name`=", t, name).Exist(d)
	if has {
		return true
	}
	var key int = 1
	//查找最大key
	has, _ = db.Where("`type`=?", t).Cols("key").Desc("key").Get(d)
	if has {
		key = d.Key + 1
	}
	r, e := db.Exec("insert into `dict` (`type`,`key`,`name`) values(?,?,?)", t, key, name)
	if e != nil {
		return false
	}
	if n, _ := r.RowsAffected(); n == 0 {
		return false
	}
	return true
}

func (d *Dicts) query(t int) ([]Dict, error) {
	if val, ok := dictListCaches.Get(t); ok {
		if _val, _ok := val.([]Dict); _ok {
			return _val, nil
		}
		return nil, nil
	}

	dictListLocker[t].Lock()
	defer dictListLocker[t].Unlock()

	if val, ok := dictListCaches.Get(t); ok {
		if _val, _ok := val.([]Dict); _ok {
			return _val, nil
		}
		return nil, nil
	}
	var ds []Dict
	err := db.Where("`type`=?", t).NoAutoCondition().Cols("key", "name").Find(&ds)
	dictListCaches.Set(t, ds)

	if err != nil {
		return nil, err
	}

	return ds, nil
}

func (d *Dicts) get(t int, key int) string {
	if val, ok := dictCaches[t].Get(key); ok {
		if _val, _ok := val.(string); _ok {
			return _val
		}
		return ""
	}

	dictLocker[t].Lock()
	defer dictLocker[t].Unlock()

	if val, ok := dictCaches[t].Get(key); ok {
		if _val, _ok := val.(string); _ok {
			return _val
		}
		return ""
	}

	db.Where("`type`=? and `key`=?", t, key).NoAutoCondition().Cols("name").Get(d)
	dictCaches[t].Set(key, d.Name)

	return d.Name
}
