package main

import (
	"fmt"
	//"./DataBase"
	"github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/sqlite"
)

type Tag struct{
	gorm.Model
	TagName string
}


func main(){
    db, err := gorm.Open("sqlite3", "Spider.db3")
    if(err!=nil){
        fmt.Println("连接数据库失败...")
    }
    
	defer db.Close()

//   if(!db.HasTable(&DataBase.Tag{})){
// 	db.CreateTable(&DataBase.Tag{})
//  }

// if(!db.HasTable(&DataBase.Website{})){
//    db.CreateTable(&DataBase.Website{})
// }
//  if(!db.HasTable(&DataBase.Article{})){
// 	db.CreateTable(&DataBase.Article{})

// }
//db.AutoMigrate(&Tag{})
db.AutoMigrate(&Tag{})


//db.Create(&Product{Code: "L1212", Price: 1000})




  fmt.Println("初始化数据库完成...")
}