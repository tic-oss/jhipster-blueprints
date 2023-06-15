package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"time"
	"github.com/google/uuid"
	"github.com/micro/micro/v3/service/errors"
	"github.com/micro/micro/v3/service/logger"
	db "<%= packageName %>/proto"
	gorm2 "<%= packageName %>/pkg"
	"google.golang.org/protobuf/types/known/structpb"
	"gorm.io/datatypes"
	<%_ if (postgresql){  _%>
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	<%_ } _%>
    <%_ if (mongodb){ _%> 
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
	<%_ } _%>
)
const idKey = "id"
const _idKey = "_id"
<%_ if (postgresql){  _%>
const stmt = "create table if not exists %v(id text not null, data jsonb, primary key(id)); alter table %v add created_at timestamptz; alter table %v add updated_at timestamptz"
<%_ } _%>

type Record struct {
	ID   string  `json:"_id" bson:"_id" `
	Data datatypes.JSON `json:"data"`
	// private field, ignored from gorm
	table     string `gorm:"-"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Db struct {
	gorm2.Helper
}

<%_ if (postgresql){  _%>
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
		if mid, ok := m[idKey].(string); ok {
			id = mid
		} else {
			id = uuid.New().String()
			m[idKey] = id
		}
	}
	bs, _ := json.Marshal(m)
	err = db.Table(tableName).Create(&Record{
		ID:   id,
		Data: bs,
	}).Error
	if err != nil {
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
<%_ } _%>

<%_ if (mongodb){ _%> 
	func GetClient() *mongo.Client {
		clientOptions := options.Client().ApplyURI("mongodb+srv://harsha:harsha@cluster0.l7oje6h.mongodb.net/?retryWrites=true&w=majority")
		client, err := mongo.NewClient(clientOptions)
		if err != nil {
			log.Fatal(err)
		}
		err = client.Connect(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		return client
	}
	
	func (e *Db) Create(ctx context.Context, req *db.CreateRequest, rsp *db.CreateResponse) error {
		if len(req.Record.AsMap()) == 0 {
			return errors.BadRequest("db.create", "missing record")
		}
		tableName :="temp"
		logger.Infof("Inserting into table '%v'", tableName)
		c := GetClient()
		collection := c.Database("temp").Collection("events")
		m := req.Record.AsMap()
		id := req.Id
		if len(id) == 0 {
			logger.Infof("hii error2")
			if mid, ok := m[_idKey].(string); ok {
				id = mid
			} else {
				id = uuid.New().String()
				m[_idKey] = id
			}
		}
		log.Println("%v",m)
		log.Println("%v",req.Record)
		bs, _ := json.Marshal(m)
		log.Printf("%v",bs)
		logger.Infof("%v",Record{ID:id,Data:bs})
		_, err := collection.InsertOne(context.TODO(),&Record{
					ID:   id,
					Data: bs,
				})
		if err != nil {
			logger.Infof("Error on inserting new event", err)
			return err
		}
		rsp.Id=id
		return nil
	}
	
	func (e *Db) Update(ctx context.Context, req *db.UpdateRequest, rsp *db.UpdateResponse) error {
		if len(req.Record.AsMap()) == 0 {
			return errors.BadRequest("db.update", "missing record")
		}
		logger.Infof("%v",req.Record)
		tableName :="temp"
		logger.Infof("Updating table '%v'", tableName)
		c := GetClient()
		collection := c.Database("temp").Collection("events")
		m := req.Record.AsMap()
		id := req.Id
		if len(id) == 0 {
				return fmt.Errorf("update failed: missing id")
		}
	
			rec := Record{}
			cle := collection.FindOne(context.TODO(), bson.M{"_id":req.Id})
			cle.Decode(&rec)
			logger.Infof("%v",rec)
			old := map[string]interface{}{}
			err := json.Unmarshal(rec.Data, &old)
			logger.Infof("%v",old)
			if err != nil {
				return err
			}
			for k, v := range m {
				logger.Infof("%v:%v",k,v)
				old[k] = v
			}
			bs, _ := json.Marshal(old)
			update := bson.D{{"$set", bson.D{{"data", bs}}}}
			collection.UpdateOne(context.TODO(),bson.M{"_id":id},update)
			return nil
	}
	
	func (e *Db) Read(ctx context.Context, req *db.ReadRequest, rsp *db.ReadResponse) error {
		recs := Record{}
		c := GetClient()
		collection := c.Database("temp").Collection("events")
		if req.Id != "" {
			logger.Infof("Query by id: %v", req.Id)
			logger.Infof("%v",Record{ID:req.Id})
			cle := collection.FindOne(context.TODO(), bson.M{"_id":req.Id})
			log.Printf("%v",cle)
			cle.Decode(&recs)
			log.Printf("%v",recs)
		} 
			rsp.Records = []*structpb.Struct{}
			m, err := recs.Data.MarshalJSON()
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
	
		return nil
	}
	
	func (e *Db) Delete(ctx context.Context, req *db.DeleteRequest, rsp *db.DeleteResponse) error {
		if len(req.Id) == 0 {
			return errors.BadRequest("db.delete", "missing id")
		}
		tableName :="temp"
		logger.Infof("Deleting from table '%v'", tableName)
		c := GetClient()
		collection := c.Database("temp").Collection("events")
		collection.DeleteOne(ctx, bson.M{"_id": req.Id})
		return nil
	}
<%_ } _%>
