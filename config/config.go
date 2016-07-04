package config

import (
  "log"
  "io/ioutil"
  "path/filepath"
  "gopkg.in/yaml.v2"
)

var App Config

// Struct for Application configuration
type Config struct{
  Database Database
}

// Struct for Database configuration
type Database struct{
  Username string
  Password string
  Database string
  Port string
}

// To initialize configuration
func InitConfig(){
  path, err := filepath.Abs("./config/application.yml")
  source, err := ioutil.ReadFile(path)
  if err != nil{
    log.Fatal(err)
  }
  err = yaml.Unmarshal(source, &App)
  if err != nil{
    log.Fatal(err)
  }
}

