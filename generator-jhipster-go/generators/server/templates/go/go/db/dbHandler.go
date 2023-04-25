package config

import (   
	"gorm.io/driver/postgres"
	"gorm.io/gorm"   
	"com.cmi.tic/domains"
	"com.cmi.tic/customlogger"
	"os"
	// ft "com.cmi.tic/fileutil"
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
    db_host :=goDotEnvVariable("db_host")
    db_user :=goDotEnvVariable("db_user")
	db_name :=goDotEnvVariable("db_name")
	db_port :=goDotEnvVariable("db_port")
	// props, _ := ft.ReadPropertiesFile("config.properties")
	dsn := "host="+db_host+" user="+db_user+" dbname="+db_name+" port="+db_port;
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})   
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