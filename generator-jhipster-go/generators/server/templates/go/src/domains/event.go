package domains

// Event represents the model for an event
type Event struct {    
	ID int `json:”ID”`   
	Title string `json:”Title”`   
	Description string `json:”Description”`
}