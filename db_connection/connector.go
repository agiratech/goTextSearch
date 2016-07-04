package db_connection

import (
  "log"
  "fmt"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/postgres"

)

// Gorp Implementation
var db *gorm.DB
var err error

func InitDb() {
  InitConfig()
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
  return fmt.Sprintf("user=%s password=%s dbname=%s port=%s", App.Database.Username, App.Database.Password, App.Database.Database, App.Database.Port)
}