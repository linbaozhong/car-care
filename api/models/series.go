package models

type Series struct {
	ID      int    `json:"id" xorm:"pk autoincr 'id'"`
	BrandID int    `json:"brand_id" xorm:"'brand_id'"`
	Name    string `json:"name" xorm:"'name'"`
	Tires   string `json:"tires" xorm:"'tires'"`
	Sort    int8   `json:"sort" xorm:"'sort'"`
}

func (*Series) TableName() string {
	return "series"
}

func NewSeries() *Series {
	return &Series{}
}

//GetAllSeries 读取全部车系
func GetAllSeries() ([]Series, error) {
	series := make([]Series, 0, 0)
	err := db.OrderBy("brand_id").Asc("sort").Find(&series)
	return series, err
}
