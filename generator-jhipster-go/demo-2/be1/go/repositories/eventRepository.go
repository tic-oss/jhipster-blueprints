package repositories
import (   
	"com.cmi.tic/db"   
	"com.cmi.tic/domains"   
	"com.cmi.tic/errors"
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

 
	func UpdateEvents (updates *map[string]interface{},id int) (*errors.HttpError){
		var updateEvent domains.Event
		if err := config.Database.First(&updateEvent, id).Error; err != nil {
			return errors.DataAccessLayerError("No field with the given id")
		}
	
		if err := config.Database.Model(&updateEvent).Updates(updates).Error; err != nil {
			return errors.DataAccessLayerError("Reverify the entities and try again")
		}
	
		return nil
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

