package models

import "fmt"

type ClerkInfo struct {
	ID   int64  `json:"id" xorm:"pk 'id'"`
	Name string `json:"name" xorm:"'name'"`
}

func NewClerkInfo() *ClerkInfo {
	return &ClerkInfo{}
}
func (*ClerkInfo) TableName() string {
	return "clerk_info"
}
func GetClerkInfo(id int64) (*ClerkInfo, error) {
	clerk := NewClerkInfo()
	has, e := db.ID(id).Get(clerk)
	fmt.Println(clerk)
	if has {
		return clerk, nil
	}
	return nil, e
}
func UpdateClerkInfo(id int64, name string) (bool, error) {
	clerk_info, e := GetClerkInfo(id)
	if clerk_info == nil {
		clerk_info = NewClerkInfo()
		clerk_info.ID = id
		clerk_info.Name = name
		n, e := db.Insert(clerk_info)
		return n > 0, e
	}
	clerk_info.Name = name
	n, e := db.ID(id).Cols("name").Update(clerk_info)
	return n > 0, e
}
