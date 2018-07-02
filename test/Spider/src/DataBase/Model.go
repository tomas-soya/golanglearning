package DataBase

import (
	"time"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Article struct{
    gorm.Model
    Id  int `gorm:"AUTO_INCREMENT"` // 自增
    Title  string
    Author string  `gorm:"size:255"`       // string默认长度为255, 使用这种tag重设。
    Content string
	ReleaseTime time.Time
	URL string
	Website Website
	Tags []Tag
}

type Website struct{
	gorm.Model
	Id int `gorm:"AUTO_INCREMENT"` // 自增
	WebsiteURL string
	WebsiteCHName string
}

type Tag struct{
	gorm.Model
	Id uint 
	TagName string
}