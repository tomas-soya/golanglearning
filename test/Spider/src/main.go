package main

import (
	"fmt"
	"./DataBase"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	//"github.com/PuerkitoBio/goquery"
	"net/http"
)

const(
	Websiteurl string="http://www.hsu.edu.cn/17/list.htm"
	IsDBInited bool=false
)

func main(){
    db, err := gorm.Open("sqlite3", "Spider.db3")
    if(err!=nil){
        fmt.Println("连接数据库失败...")
    }
    
	defer db.Close()

db.AutoMigrate(&DataBase.Tag{},&DataBase.Website{},&DataBase.Article{})
  fmt.Println("初始化数据库完成...")
}

func getResponse(url string) *http.Response {
    client := &http.Client{}
    request, _ := http.NewRequest("GET", url, nil)
    request.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 6.1; Win64; x64; rv:50.0) Gecko/20100101 Firefox/50.0")
	response, _ := client.Do(request)
    return response
}