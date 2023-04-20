package config

import (   
	"gorm.io/driver/postgres"
	"gorm.io/gorm"   
	"<%= packageName %>/domains"
	"<%= packageName %>/customlogger"
	ft "<%= packageName %>/fileutil"
)
	
var Database *gorm.DB

func DbConnect(){   
	props, _ := ft.ReadPropertiesFile("config.properties")
	dsn := "host="+props["db_host"]+" user="+props["db_user"]+" dbname="+props["db_name"]+" port="+props["db_port"];
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