package DataBase

// import (
//     "github.com/jinzhu/gorm"
//     _ "github.com/jinzhu/gorm/dialects/sqlite"
// )



// func dbConnection(){
//     db, err := gorm.Open("sqlite3", "C:\\sqlite\\DataBase\\Spider.db3")
//     if(err!=nil){
//         //return nil
//     }
//     db.AutoMigrate(&Article{})
//     db.AutoMigrate(&Website{})
//     db.AutoMigrate(&Tag{})
//   defer db.Close()
// }