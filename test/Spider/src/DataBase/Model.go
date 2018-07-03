package DataBase

import (
	"time"
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Article struct{
    gorm.Model
    Title  string
    Author string  `gorm:"size:255"`       // string默认长度为255, 使用这种tag重设。
    Content string
	ReleaseTime time.Time
	URL string
	//Website Website `gorm:"ForeignKey:WebsiteID"`
	//WebsiteID int
	//Tags []Tag
}

type Website struct{
	gorm.Model
	WebsiteURL string
	WebsiteCHName string
}

type Tag struct{
	gorm.Model
	TagName string
}