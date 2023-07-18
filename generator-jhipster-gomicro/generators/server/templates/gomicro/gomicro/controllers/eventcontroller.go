package controllers

import (
	auth "<%= packageName %>/auth"
	"github.com/gorilla/mux"
	"net/http"
	"encoding/json"
	<%_ if (restServer?.length){  _%>
	eureka "<%= packageName %>/eurekaregistry"
	<%_ } _%>
	<%_ if (postgresql||mongodb){  _%>
	"<%= packageName %>/handler"
	<%_ } _%>
)

type EventController struct {
}

func (t EventController) RegisterRoutes(r *mux.Router) {
	<%_ if (postgresql||mongodb){  _%>
	r.Handle("/event",auth.Protect(http.HandlerFunc(handler.AddEvent))).Methods(http.MethodPost)
	r.Handle("/events",auth.Protect(http.HandlerFunc(handler.GetEvents))).Methods(http.MethodGet)
	r.Handle("/events/{id}",auth.Protect(http.HandlerFunc(handler.ReadEventById))).Methods(http.MethodGet)
	r.Handle("/update",auth.Protect(http.HandlerFunc(handler.UpdateEvent))).Methods(http.MethodPatch)
	r.Handle("/delete/{id}",auth.Protect(http.HandlerFunc(handler.DeleteEvent))).Methods(http.MethodDelete)
	<%_ } _%>
	r.HandleFunc("/management/health/readiness", func(w http.ResponseWriter, _ *http.Request) {   
	json.NewEncoder(w).Encode(map[string]interface{}{"status": "UP","components":map[string]interface{} {"readinessState": map[string]interface{}{"status": "UP"}}})}).Methods(http.MethodGet)
	<%_ if (restClient){  _%>
	r.HandleFunc("/api/services/<%= baseName %>",func(w http.ResponseWriter, _ *http.Request){
	logger.Infof("response sent")
	json.NewEncoder(w).Encode(map[string]string{"server": "UP"})
	}).Methods(http.MethodGet)
	<%_ } _%>	
	<%_ if (restServer?.length){  for(var i=0;i<restServer.length;i++){_%>
	r.HandleFunc("/api/services/<%= restServer[i] %>",func(w http.ResponseWriter, r *http.Request) { eureka.Client(w,r,"<%= restServer[i] %>")}).Methods(http.MethodGet)
	<%_ }} _%>	
	r.HandleFunc("/hello",func(w http.ResponseWriter, _ *http.Request){
	json.NewEncoder(w).Encode("helloworld")}).Methods(http.MethodGet)
	r.HandleFunc("/management/health/liveness", func(w http.ResponseWriter, _ *http.Request) {     
	json.NewEncoder(w).Encode(map[string]interface{}{"status": "UP","components":map[string]interface{} {"livenessState": map[string]interface{}{"status": "UP"}}})}).Methods(http.MethodGet)
}