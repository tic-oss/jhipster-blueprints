package config 

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm" 
	"os"
	"log"
	"fmt"
)

var stmt = "create table if not exists %v(id text not null, title text,description text, primary key(id));"
var tableName="event"

func GetClient() *gorm.DB {
	log.Print("1")
    db_url :=os.Getenv("DB_URL")
	db, err := gorm.Open(postgres.Open(db_url), &gorm.Config{})   
	if err != nil{      
		panic("failed to connect database")
		log.Print("failed to connect database")  	
	}   
	log.Print("Database connected")
	db.Exec(fmt.Sprintf(stmt, tableName)) 
	return db
}