package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"gorm.io/driver/postgres"
	"github.com/google/uuid"
	"github.com/micro/micro/v3/service/errors"
	"github.com/micro/micro/v3/service/logger"
	db "<%= packageName %>/proto"
	gorm2 "<%= packageName %>/pkg"
	"google.golang.org/protobuf/types/known/structpb"
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

const idKey = "id"
const _idKey = "_id"
const stmt = "create table if not exists %v(id text not null, data jsonb, primary key(id)); alter table %v add created_at timestamptz; alter table %v add updated_at timestamptz"

type Record struct {
	ID   string
	Data datatypes.JSON `json:"data"`
	// private field, ignored from gorm
	table     string `gorm:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Db struct {
	gorm2.Helper
}


func (e *Db) Create(ctx context.Context, req *db.CreateRequest, rsp *db.CreateResponse) error {
	if len(req.Record.AsMap()) == 0 {
		return errors.BadRequest("db.create", "missing record")
	}
	tableName :="temp"
	logger.Infof("Inserting into table '%v'", tableName)
	db, err := gorm.Open(postgres.Open("postgresql://go@localhost:5433/postgres"), &gorm.Config{})   
	if err != nil {
		logger.Infof("hii error")
		return err
	}
	db.Exec(fmt.Sprintf(stmt, tableName, tableName, tableName))
	m := req.Record.AsMap()
	id := req.Id
	if len(id) == 0 {
		logger.Infof("hii error2")
		if mid, ok := m[idKey].(string); ok {
			id = mid
		} else {
			id = uuid.New().String()
			m[idKey] = id
		}
	}
	bs, _ := json.Marshal(m)
	logger.Infof("hii error3")
	err = db.Table(tableName).Create(&Record{
		ID:   id,
		Data: bs,
	}).Error
	if err != nil {
		logger.Infof("hii error4")
		return err
	}

	// set the response id
	rsp.Id = id
	logger.Infof("hii error5")

	return nil
}

func (e *Db) Update(ctx context.Context, req *db.UpdateRequest, rsp *db.UpdateResponse) error {
	if len(req.Record.AsMap()) == 0 {
		return errors.BadRequest("db.update", "missing record")
	}
	tableName :="temp"
	logger.Infof("Updating table '%v'", tableName)
	db, err := gorm.Open(postgres.Open("postgresql://go@localhost:5433/postgres"), &gorm.Config{})   
	if err != nil {
		return err
	}
	m := req.Record.AsMap()

	id := req.Id
	if len(id) == 0 {
		var ok bool
		id, ok = m[idKey].(string)
		if !ok {
			return fmt.Errorf("update failed: missing id")
		}
	}

	return db.Transaction(func(tx *gorm.DB) error {
		rec := []Record{}
		err = tx.Table(tableName).Where("id = ?", id).Find(&rec).Error
		if err != nil {
			return err
		}
		if len(rec) == 0 {
			return fmt.Errorf("update failed: not found")
		}
		old := map[string]interface{}{}
		err = json.Unmarshal(rec[0].Data, &old)
		if err != nil {
			return err
		}
		for k, v := range m {
			old[k] = v
		}
		bs, _ := json.Marshal(old)

		return tx.Table(tableName).Save(&Record{
			ID:   id,
			Data: bs,
		}).Error
	})
}

func (e *Db) Read(ctx context.Context, req *db.ReadRequest, rsp *db.ReadResponse) error {
	recs := []Record{}
    tableName :="temp"
	db, err := gorm.Open(postgres.Open("postgresql://go@localhost:5433/postgres"), &gorm.Config{})   
	if err != nil {
		return err
	}
	db = db.Table(tableName)
	if req.Id != "" {
		logger.Infof("Query by id: %v", req.Id)
		db = db.Where("id = ?", req.Id)
	} 
	err = db.Debug().Find(&recs).Error
	if err != nil {
		return err
	}

	rsp.Records = []*structpb.Struct{}
	for _, rec := range recs {
		m, err := rec.Data.MarshalJSON()
		if err != nil {
			return err
		}

		ma := map[string]interface{}{}
		json.Unmarshal(m, &ma)

		m, _ = json.Marshal(ma)
		s := &structpb.Struct{}

		if err = s.UnmarshalJSON(m); err != nil {
			return err
		}
		rsp.Records = append(rsp.Records, s)
	}

	return nil
}

func (e *Db) Delete(ctx context.Context, req *db.DeleteRequest, rsp *db.DeleteResponse) error {
	if len(req.Id) == 0 {
		return errors.BadRequest("db.delete", "missing id")
	}
	tableName :="temp"
	logger.Infof("Deleting from table '%v'", tableName)
	db, err := gorm.Open(postgres.Open("postgresql://go@localhost:5433/postgres"), &gorm.Config{})   
	if err != nil {
		return err
	}
	return db.Table(tableName).Delete(Record{
		ID: req.Id,
	}).Error
}