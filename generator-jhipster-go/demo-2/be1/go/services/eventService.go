package services

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"com.cmi.tic/domains"
	"com.cmi.tic/errors"
	"com.cmi.tic/repositories"
	"com.cmi.tic/customlogger"
	"strconv"
	"log"
)

// Event represents the model for an event
type Event struct {    
	ID int `json:”ID”`   
	Title string `json:”Title”`   
	Description string `json:”Description”`
}

type Response struct {
    Status string `json:"status"`
}


func Health(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"hello");
}


func Liveness(w http.ResponseWriter, r *http.Request) {
    response := Response{Status: "UP"}

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}

func Readiness(w http.ResponseWriter, r *http.Request){
    response := Response{Status: "UP"}

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(response)
}

// CreateEvent godoc
// @Summary Create a new event
// @Description Create a new event with the input paylod
// @Tags events
// @Accept  json
// @Produce  json
// @Param event body Event true "Create event"
// @Success 200 {object} Event
// @Router /event [post]
func CreateEvent(w http.ResponseWriter, r *http.Request) {
	var newEvent domains.Event
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		customlogger.Printfun("error","Kindly enter data with the event title and description only in order to update")    
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)
    
	ev, httpErr := repositories.SaveEvent(&newEvent)
	if httpErr != nil {
		w.WriteHeader(httpErr.Code)
		json.NewEncoder(w).Encode(errors.UnauthorizedError())
		customlogger.Printfun("error","Unauthorized Error")
		return
	}

	w.WriteHeader(http.StatusCreated)
	customlogger.Printfun("info","Event created with Id:"+strconv.Itoa(newEvent.ID))    
	json.NewEncoder(w).Encode(&ev)
}

// GetEvent godoc
// @Summary Fetch an event By Id
// @Description Fetch the event details by id
// @Tags events
// @Accept  json
// @Produce  json
// @Param id header int true "Get event"
// @Success 200 {object} Event
// @Router /event/{id} [get]
func GetOneEvent(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	idStr := params[ "id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(
			errors.BadRequestError("Id must be an integer"))
			customlogger.Printfun("error","Id must be an integer")    
		return
	}

	event := repositories.FindOneEventById(id)

	if event == nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(errors.NotFoundError())
		customlogger.Printfun("error","Event Not found")    
		return
	}
	customlogger.Printfun("info","Fetched event with Id="+idStr)    
	json.NewEncoder(w).Encode(&event)
}

// UpdateEvent godoc
// @Summary Update an event
// @Description Updates an event with the input paylod
// @Tags events
// @Accept  json
// @Produce  json
// @Param body {object} true "Update event"
// @Param id header int true "Update event"
// @Router /update [patch]
func UpdateEvent(w http.ResponseWriter, r *http.Request){
	params := mux.Vars(r)
	idStr := params[ "id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(
			errors.BadRequestError("Id must be an integer"))
			customlogger.Printfun("error","Id must be an integer")  
		return
	}
	var unknownMap map[string]interface{}
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
		customlogger.Printfun("info","Kindly enter data with the event title and description only in order to update")  
	}
	err = json.Unmarshal([]byte(reqBody), &unknownMap)
	if err != nil {
		log.Fatal(err)
	}
    
	httpErr := repositories.UpdateEvents(&unknownMap,id)
	
	if httpErr != nil {
		w.WriteHeader(httpErr.Code)
		json.NewEncoder(w).Encode(errors.UnauthorizedError())
		customlogger.Printfun("error","Unauthorized Error")  
		return
	}

	w.WriteHeader(http.StatusCreated)
	customlogger.Printfun("info","Event updated with Id:"+idStr)  
}

// DeleteEvent godoc
// @Summary Delete an event By Id
// @Description Delete the event by id
// @Tags events
// @Accept  json
// @Produce  json
// @Param id header int true "Delete event"
// @Router /delete/{id} [delete]
func DeleteEvent(w http.ResponseWriter,r *http.Request){
	params := mux.Vars(r)
	idStr := params[ "id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(
			errors.BadRequestError("Id must be an integer"))
			customlogger.Printfun("error","Id must be an integer")  	
		return
	}
	_,errr := repositories.DeleteEventById(id)
  
	if errr != nil {
		json.NewEncoder(w).Encode(errr)
		customlogger.Printfun("error",errr.Error)  	
		return
	}
	customlogger.Printfun("info","Deleted event with Id="+idStr)  	
}

// GetEvents godoc
// @Summary Get details of all events
// @Description Get details of all events
// @Tags events
// @Accept  json
// @Produce  json
// @Success 200 {array} Event
// @Router /events [get]
func AllEvents(w http.ResponseWriter, r *http.Request) {
	events := repositories.FindAllEvents()
	customlogger.Printfun("info","Fetched all events")	
	json.NewEncoder(w).Encode(&events)
}