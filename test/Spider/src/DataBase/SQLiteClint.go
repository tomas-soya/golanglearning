package DataBase

import (
    "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"fmt"
)

	var IsDBInited bool=false

func dbConnection() *gorm.DB {
    db, err := gorm.Open("sqlite3", "Spider.db3")
    if(err!=nil){
        fmt.Println("连接数据库失败...")
    }
    
	defer db.Close()
if(IsDBInited==false){
	db.AutoMigrate(&Tag{},Website{},Article{})
	fmt.Println("初始化数据库完成...")
	IsDBInited=true
}else{
	fmt.Println("数据库已经初始化...")
}

  return db
}