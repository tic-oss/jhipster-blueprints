package handler

import (
	"context"
	"encoding/json"
	"net/http"
	pb "<%= packageName %>/proto"
	config "<%= packageName %>/db"  
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/micro/micro/v3/service/logger"
)

var instance *mongo.Client
var event pb.Event

func InitializeMongoDb(){
	instance = config.GetInstance()
}

func AddEvent(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	_ = json.NewDecoder(request.Body).Decode(&event)
	logger.Infof("%v",event)
	collection := instance.Database("temp").Collection("events")
	result, err := collection.InsertOne(context.Background(), event)
	if err != nil {
		logger.Errorf(err.Error())
		response.WriteHeader(400)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	} else {
		logger.Infof("Inserted object with Id:" + event.Id)
		json.NewEncoder(response).Encode(result)
	}
}

func ReadEventById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id := params["id"]
	collection := instance.Database("temp").Collection("events")
    err := collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&event)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		logger.Errorf(err.Error())
		return
	}
	logger.Infof("Fetched event by ID:" + id)
	json.NewEncoder(response).Encode(event)
}

func GetEvents(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var events []pb.Event
	collection := instance.Database("temp").Collection("events")
	cursor, err := collection.Find(context.Background(), bson.M{})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		logger.Errorf(err.Error())
		return
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		cursor.Decode(&event)
		events = append(events, event)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		logger.Errorf(err.Error())
		return
	}
	logger.Infof("Fetched all events")
	json.NewEncoder(response).Encode(events)
}

func UpdateEvent(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	collection := instance.Database("temp").Collection("events")
	_ = json.NewDecoder(request.Body).Decode(&event)
	filter := bson.D{{Key: "id", Value: event.Id}}
	result, err := collection.ReplaceOne(context.Background(), filter, event)
	if err != nil {
		logger.Errorf(err.Error())
		panic(err)
	}
	logger.Infof("Updated Event With Id:" + event.Id)
	json.NewEncoder(response).Encode(result)
}

func DeleteEvent(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	id := params["id"]
	collection := instance.Database("temp").Collection("events")
	filter := bson.D{{Key: "id", Value: id}}
	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		logger.Errorf(err.Error())
		panic(err)
	}
	logger.Infof("Deleted Event With Id:" + id)
	json.NewEncoder(response).Encode(result)
}