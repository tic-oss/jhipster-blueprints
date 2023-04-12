package controllers

import (
	"net/http"
	"<%= packageName %>/src/services"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

type EventController struct {
}

func (t EventController) RegisterRoutes(router *mux.Router) {
	router.Handle("/",http.HandlerFunc(services.Login)).Methods("GET")
	router.Handle("/redirect",http.HandlerFunc(services.Redirect)).Methods("GET")
	// create
	router.Handle("/event", services.Protect(http.HandlerFunc(services.CreateEvent))).Methods("POST")
	//read
	router.Handle("/events/{id}", services.Protect(http.HandlerFunc(services.GetOneEvent))).Methods("GET")
	//read-all
	router.Handle("/events", http.HandlerFunc(services.AllEvents)).Methods("GET")
	router.Handle("/healthcheck", http.HandlerFunc(services.Health)).Methods("GET")
	//update
	router.Handle("/update",http.HandlerFunc(services.UpdateEvent)).Methods("POST")
	//delete
	router.Handle("/delete/{id}",http.HandlerFunc(services.DeleteEvent)).Methods("GET")
	// Swagger
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
}