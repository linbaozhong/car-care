package models

type Clerk struct {
	ID         int64  `json:"id" xorm:"pk autoincr 'id'"`
	Mobile     int    `json:"mobile" xorm:"'mobile'"`
	Pwd        string `json:"pwd" xorm:"'pwd'"`
	Auto       string `json:"auto" xorm:"'auto'"`
	StoreID    int64  `json:"store_id" xorm:"'store_id'"`
	RoleID     int8   `json:"role_id" xorm:"'role_id'"`
	CreateTime int64  `json:"create_time" xorm:"'create_time'"`
}

func (*Clerk) TableName() string {
	return "clerk"
}

func NewClerk() *Clerk {
	return &Clerk{}
}

func AddClerk(clerk *Clerk) (bool, error) {
	n, e := db.Insert(clerk)
	return n > 0, e
}
func GetClerk(id int64) (*Clerk, error) {
	clerk := NewClerk()
	has, e := db.ID(id).Get(clerk)
	if has {
		return clerk, nil
	}
	return nil, e
}
func GetClerkByMobile(mobile int) (*Clerk, error) {
	clerk := NewClerk()
	has, e := db.Where("`mobile`=?", mobile).Get(clerk)
	if has {
		return clerk, nil
	}
	return nil, e
}

