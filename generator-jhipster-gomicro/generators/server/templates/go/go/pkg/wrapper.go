package gorm

import (
	"database/sql"
	"sync"
	"gorm.io/gorm"
)

type Helper struct {
	sync.RWMutex
	gormConns  map[string]*gorm.DB
	dbConn     *sql.DB
	migrations []interface{}
}

func (h *Helper) Migrations(migrations ...interface{}) *Helper {
	h.migrations = migrations
	return h
}

func (h *Helper) DBConn(conn *sql.DB) *Helper {
	h.dbConn = conn
	h.gormConns = map[string]*gorm.DB{}
	return h
}

