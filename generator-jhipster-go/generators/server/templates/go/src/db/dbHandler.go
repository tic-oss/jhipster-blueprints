package config

import (   
	"gorm.io/driver/postgres"
	"gorm.io/gorm"   
	"<%= packageName %>/src/domains"
	"<%= packageName %>/src/customlogger"
	ft "<%= packageName %>/fileutil"
)
	
var Database *gorm.DB

func DbConnect(){   
	props, _ := ft.ReadPropertiesFile("config.properties")
	logger := customlogger.GetInstance()
	dsn := "host="+props["db_host"]+" user="+props["db_user"]+" dbname="+props["db_name"]+" port="+props["db_port"];
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})   
	Database = db   
	if err != nil{      
		panic("failed to connect database")
		logger.ErrorLogger.Println("failed to connect database")   
	}   
	runMigrations()
	logger.InfoLogger.Println("Database connected")
}
	
func runMigrations(){   
  Database.AutoMigrate(&domains.Event{})
}