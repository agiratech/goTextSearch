package data_groups

import (
  "log"
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"
  "github.com/agiratech/goTextSearch/config"

)

// Gorp Implementation
var db *gorm.DB
var err error

func InitDb() {
  db, err = gorm.Open("postgres", getFormat())
  checkErr(err, "sql.Open failed")
  fmt.Println("db connected",db)
}

 func checkErr(err error, msg string) {
    if err != nil {
        log.Fatalln(msg, err)
    }
}

func getFormat()string{
  return fmt.Sprintf("user=%s password=%s dbname=%s port=%s", config.App.Database.Username, config.App.Database.Password, config.App.Database.Database, config.App.Database.Port)
}