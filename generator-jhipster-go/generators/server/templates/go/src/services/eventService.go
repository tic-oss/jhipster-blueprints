package services

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"<%= packageName %>/src/domains"
	"<%= packageName %>/src/errors"
	"<%= packageName %>/src/repositories"
	"<%= packageName %>/src/customlogger"
	"strconv"
)

// Event represents the model for an event
type Event struct {    
	ID int `json:”ID”`   
	Title string `json:”Title”`   
	Description string `json:”Description”`
}


func Health(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"hello");
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
	logger := customlogger.GetInstance()
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)
    
	ev, httpErr := repositories.SaveEvent(&newEvent)
	if httpErr != nil {
		w.WriteHeader(httpErr.Code)
		json.NewEncoder(w).Encode(errors.UnauthorizedError())
		logger.ErrorLogger.Println(errors.UnauthorizedError())
		return
	}

	w.WriteHeader(http.StatusCreated)
	logger.InfoLogger.Println("Event created with Id:"+strconv.Itoa(newEvent.ID))
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
	logger := customlogger.GetInstance()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(
			errors.BadRequestError("Id must be an integer"))
		logger.ErrorLogger.Println(errors.BadRequestError("Id must be an integer"))
		return
	}

	event := repositories.FindOneEventById(id)

	if event == nil {
		w.WriteHeader(404)
		json.NewEncoder(w).Encode(errors.NotFoundError())
		logger.ErrorLogger.Println(errors.NotFoundError())

		return
	}
    logger.InfoLogger.Println("Fetched event with Id="+idStr)
	json.NewEncoder(w).Encode(&event)
}

// UpdateEvent godoc
// @Summary Update an event
// @Description Updates an event with the input paylod
// @Tags events
// @Accept  json
// @Produce  json
// @Param event body Event true "Update event"
// @Success 200 {object} Event
// @Router /update [post]
func UpdateEvent(w http.ResponseWriter, r *http.Request){
	var newEvent domains.Event
	reqBody, err := ioutil.ReadAll(r.Body)
	logger := customlogger.GetInstance()
	if err != nil {
		fmt.Fprintf(w, "Kindly enter data with the event title and description only in order to update")
	}

	json.Unmarshal(reqBody, &newEvent)
    
	ev, httpErr := repositories.UpdateEvents(&newEvent)
	
	if httpErr != nil {
		w.WriteHeader(httpErr.Code)
		json.NewEncoder(w).Encode(errors.UnauthorizedError())
		logger.ErrorLogger.Println(errors.UnauthorizedError())
		return
	}

	w.WriteHeader(http.StatusCreated)
	logger.InfoLogger.Println("Event updated with Id:"+strconv.Itoa(newEvent.ID))
	json.NewEncoder(w).Encode(&ev)
}

// DeleteEvent godoc
// @Summary Delete an event By Id
// @Description Delete the event by id
// @Tags events
// @Accept  json
// @Produce  json
// @Param id header int true "Delete event"
// @Success 200 {object} Event
// @Router /delete/{id} [get]
func DeleteEvent(w http.ResponseWriter,r *http.Request){
	params := mux.Vars(r)
	idStr := params[ "id"]
	logger := customlogger.GetInstance()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(400)
		json.NewEncoder(w).Encode(
			errors.BadRequestError("Id must be an integer"))
		logger.ErrorLogger.Println(errors.BadRequestError("Id must be an integer"))
		return
	}
	_,errr := repositories.DeleteEventById(id)
  
	if errr != nil {
		json.NewEncoder(w).Encode(errr)
		logger.ErrorLogger.Println(errr)
		return
	}
    logger.InfoLogger.Println("Deleted event with Id="+idStr)
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
	logger := customlogger.GetInstance()
	events := repositories.FindAllEvents()
	logger.InfoLogger.Println("Fetched all events");
	json.NewEncoder(w).Encode(&events)
}