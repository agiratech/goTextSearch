package config

import (
  "log"
  "fmt"
  "database/sql"
_ "github.com/lib/pq"
  "gopkg.in/gorp.v1"
)

// Gorp Implementation
var Dbmap *gorp.DbMap
var Db *sql.DB
var err error

func InitDb() {
  Db, err = sql.Open("postgres", getFormat())
  checkErr(err, "sql.Open failed")
  fmt.Println("db connected",Db)
  Dbmap = &gorp.DbMap{Db: Db, Dialect: gorp.PostgresDialect{}}
}

 func checkErr(err error, msg string) {
    if err != nil {
        log.Fatalln(msg, err)
    }
}

func getFormat()string{
  return fmt.Sprintf("user=%s password=%s dbname=%s port=%s", App.Database.Username, App.Database.Password, App.Database.Database, App.Database.Port)
}