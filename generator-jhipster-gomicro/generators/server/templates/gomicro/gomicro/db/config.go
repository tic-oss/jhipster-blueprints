package config 

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm" 
	"os"
	"github.com/micro/micro/v3/service/logger"
	"fmt"
)

var stmt = "create table if not exists %v(id text not null, title text,description text, primary key(id));"
var tableName="event"

func GetClient() *gorm.DB {
    db_url :=os.Getenv("GO_MICRO_DB_URL")
	db, err := gorm.Open(postgres.Open(db_url), &gorm.Config{})   
	if err != nil{      
		panic("failed to connect database")
		logger.Errorf("failed to connect database")  	
	}   
	logger.Infof("Database connected")
	db.Exec(fmt.Sprintf(stmt, tableName)) 
	return db
}