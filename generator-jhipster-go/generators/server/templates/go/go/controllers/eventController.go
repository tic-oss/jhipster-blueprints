package controllers

import (
	"net/http"
	"<%= packageName %>/services"
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
	router.Handle("/update/{id}",http.HandlerFunc(services.UpdateEvent)).Methods("PATCH")
	//delete
	router.Handle("/delete/{id}",http.HandlerFunc(services.DeleteEvent)).Methods("DELETE")
	// Swagger
	router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)


	router.Handle("/management/health/readiness", http.HandlerFunc(services.Readiness)).Methods("GET")
	router.Handle("/management/health/liveness", http.HandlerFunc(services.Liveness)).Methods("GET")
}