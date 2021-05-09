package model

import "time"

type User struct {
	Id         int64     `gorm:"column:id"`
	Name       string    `gorm:"column:name"`
	Password   string    `gorm:"column:password"`
	Phone      string    `gorm:"column:phone"`
	WxId       string    `gorm:"column:wx_id"`
	CreateTime time.Time `gorm:"column:create_time"`
	ModifyTime time.Time `gorm:"column:modify_time"`
	Role       int64     `gorm:"column:role"`
	Status     int64     `gorm:"column:status"`
	Data       string    `gorm:"column:data"`
}

func (User) TableName() string {
	return "blog_user"
}
