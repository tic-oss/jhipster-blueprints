package gorm

import (
	<%_ if (mongodb){ _%>    
    "go.mongodb.org/mongo-driver/mongo"
	<%_ } _%>
	<%_ if (postgresql){  _%>
	"database/sql"
	<%_ } _%>

)

type Helper struct {
	<%_ if (postgresql){  _%>
	dbConn     *sql.DB
	<%_ } _%>
	<%_ if (mongodb){ _%> 
	dbConn     *mongo.Client
	<%_ } _%>

}

<%_ if (postgresql){  _%>
func (h *Helper) DBConn(conn *sql.DB) *Helper {
	h.dbConn = conn
	return h
}
<%_ } _%>

<%_ if (mongodb){ _%> 
func (h *Helper) DBConn(conn *mongo.Client) *Helper {
	h.dbConn = conn
	return h
}
<%_ } _%>
	

