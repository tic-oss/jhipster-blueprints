package repositories
import (   
	"<%= packageName %>/src/db"   
	"<%= packageName %>/src/domains"   
	"<%= packageName %>/src/errors"
	// "errors"
	)

	func SaveEvent(event *domains.Event) (*domains.Event, *errors.HttpError){   
	e := config.Database.Create(&event)   
	if e.Error != nil{      
		return nil, errors.DataAccessLayerError(e.Error.Error())   
	}   
	return event, nil
	}
	

	func FindOneEventById(id int) *domains.Event{   
	var event domains.Event   
	config.Database.First(&event, id)   
	return &event
    }

 
	func UpdateEvents (event *domains.Event) (*domains.Event, *errors.HttpError){
	var updateEvent domains.Event
	result := config.Database.Model(&updateEvent).Where("id = ?", event.ID).Updates(event)
	if result.RowsAffected == 0 {
        return &updateEvent, errors.DataAccessLayerError("error")
    }
    return &updateEvent, nil
    }

	func DeleteEventById(id int) (int64, *errors.HttpError){
	var deletedEvent domains.Event
	result := config.Database.Where("id = ?", id).Delete(&deletedEvent)
    if result.RowsAffected == 0 {
        return 0, errors.DataAccessLayerError("unable to delete.Please verify")
    }
    return result.RowsAffected, nil
    }

 
	func FindAllEvents() []domains.Event {   
	var events []domains.Event   
	config.Database.Find(&events)   
	return events
    }

