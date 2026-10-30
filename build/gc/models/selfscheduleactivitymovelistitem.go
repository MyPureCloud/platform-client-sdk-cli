package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SelfscheduleactivitymovelistitemMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SelfscheduleactivitymovelistitemDud struct { 
    


    


    


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Selfscheduleactivitymovelistitem
type Selfscheduleactivitymovelistitem struct { 
    // Id - The globally unique identifier for the object.
    Id string `json:"id"`


    // Schedule - The schedule to which the shift and activity are associated
    Schedule Buschedulereference `json:"schedule"`


    // ShiftId - The ID of the shift to which this activity is associated
    ShiftId string `json:"shiftId"`


    // State - The state of the self-schedule operation
    State string `json:"state"`


    // WorkflowStatus - The processing status of this self-schedule operation
    WorkflowStatus string `json:"workflowStatus"`


    // StartDate - The original start date of this activity in ISO-8601 format
    StartDate time.Time `json:"startDate"`


    // NewDate - The new start date of this activity in ISO-8601 format
    NewDate time.Time `json:"newDate"`


    // LengthMinutes - The length of this activity in minutes
    LengthMinutes int `json:"lengthMinutes"`


    // ActivityCode - The activity code related to this shift
    ActivityCode Activitycodereference `json:"activityCode"`


    

}

// String returns a JSON representation of the model
func (o *Selfscheduleactivitymovelistitem) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Selfscheduleactivitymovelistitem) MarshalJSON() ([]byte, error) {
    type Alias Selfscheduleactivitymovelistitem

    if SelfscheduleactivitymovelistitemMarshalled {
        return []byte("{}"), nil
    }
    SelfscheduleactivitymovelistitemMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Schedule Buschedulereference `json:"schedule"`
        
        ShiftId string `json:"shiftId"`
        
        State string `json:"state"`
        
        WorkflowStatus string `json:"workflowStatus"`
        
        StartDate time.Time `json:"startDate"`
        
        NewDate time.Time `json:"newDate"`
        
        LengthMinutes int `json:"lengthMinutes"`
        
        ActivityCode Activitycodereference `json:"activityCode"`
        *Alias
    }{

        


        


        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

