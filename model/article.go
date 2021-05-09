package model

import "time"

type Article struct {
	Id         int64     `gorm:"column:id"`
	Title      string    `gorm:"column:title"`
	Content    string    `gorm:"column:content"`
	Status     int64     `gorm:"column:status"`
	CreateTime time.Time `gorm:"column:create_time"`
	ModifyTime time.Time `gorm:"column:modify_time"`
	Data       string    `gorm:"column:data"`
	UserId     int64     `gorm:"column:user_id"`
}

func (Article) TableName() string {
	return "blog_article"
}
