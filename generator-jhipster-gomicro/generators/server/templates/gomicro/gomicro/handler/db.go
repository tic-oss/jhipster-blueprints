package handler

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
	pb "<%= packageName %>/proto"
	"gorm.io/gorm" 
	config "<%= packageName %>/db"  
	"github.com/micro/micro/v3/service/logger"
)

var tableName="event"
var client *gorm.DB

func InitializeDb(){
	client=config.GetClient()
}

func CreateHandler(response http.ResponseWriter, request *http.Request){
	response.Header().Set("content-type", "application/json")
	var event *pb.Event
	_ = json.NewDecoder(request.Body).Decode(&event)
	e := client.Table(tableName).Create(&event)
	if e.Error != nil{     
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + e.Error.Error() + `" }`)) 
		return
	}   
	logger.Infof("Inserted data with Id:"+event.Id)
	json.NewEncoder(response).Encode(event)

}

func ReadByIdHandler(response http.ResponseWriter, r *http.Request){
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(r)
	id := params["id"]
	var event *pb.Event
	e := client.Table(tableName).First(&event,id)	
	if e.Error != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + e.Error.Error() + `" }`))
		return
	}
	logger.Infof("Fetched Event With Id:"+id)
	json.NewEncoder(response).Encode(event)
}

func ReadHandler(response http.ResponseWriter, r *http.Request){
	response.Header().Set("content-type", "application/json")
	var events []*pb.Event
	e := client.Table(tableName).Find(&events)
	if e.Error != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + e.Error.Error() + `" }`))
		return
	}
	logger.Infof("Fetched all Events")
	json.NewEncoder(response).Encode(events)
}

func UpdateHandler(response http.ResponseWriter, request *http.Request){
	response.Header().Set("content-type", "application/json")
	var event *pb.Event
	var ev *pb.Event
	_ = json.NewDecoder(request.Body).Decode(&event)
	client.Table(tableName).First(&ev, event.Id)
	client.Table(tableName).Model(&ev).Updates(event)
	logger.Infof("Updated Event with Id:"+event.Id)
	json.NewEncoder(response).Encode("Updated Event with Id:"+event.Id)
}

func DeleteHandler(response http.ResponseWriter, request *http.Request){
	response.Header().Set("content-type", "application/json")
	params := mux.Vars(request)
	var event *pb.Event
	id := params["id"]
	result :=client.Table(tableName).Where("id = ?", id).Delete(&event)
    if result.RowsAffected == 0 {
        response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`Unable to delete Please Verify`))
        return 
	}
	logger.Infof("Deleted Event with Id:"+event.Id)
	json.NewEncoder(response).Encode(result.RowsAffected)
}