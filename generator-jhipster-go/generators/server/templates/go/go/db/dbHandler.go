package config

import (   
	"gorm.io/driver/postgres"
	"gorm.io/gorm"   
	"<%= packageName %>/domains"
	"<%= packageName %>/customlogger"
	"os"
	"log"
	"github.com/joho/godotenv"
)
	
var Database *gorm.DB

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
	  log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
  }

func DbConnect(){   
    db_url :=goDotEnvVariable("DB_URL")
	db, err := gorm.Open(postgres.Open(db_url), &gorm.Config{})   
	Database = db   
	if err != nil{      
		panic("failed to connect database")
		customlogger.Printfun("error","failed to connect database")  	
	}   
	runMigrations()
	customlogger.Printfun("info","Database connected")  	
}
	
func runMigrations(){   
  Database.AutoMigrate(&domains.Event{})
}