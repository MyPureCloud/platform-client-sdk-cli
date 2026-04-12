package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScheduledtriggerMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScheduledtriggerDud struct { 
    Id string `json:"id"`


    


    


    


    


    


    


    SelfUri string `json:"selfUri"`

}

// Scheduledtrigger - Defines a process automation scheduled trigger.
type Scheduledtrigger struct { 
    


    // Name - The name of the scheduled trigger. Can be up to 162 characters in length.
    Name string `json:"name"`


    // Target - The target to invoke when the scheduled trigger fires
    Target Triggertarget `json:"target"`


    // Version - Version of this scheduled trigger
    Version int `json:"version"`


    // Enabled - Whether or not the scheduled trigger is enabled
    Enabled bool `json:"enabled"`


    // Schedule - The schedule configuration for when this trigger should fire
    Schedule Triggerschedule `json:"schedule"`


    // Description - Description of the trigger. Can be up to 512 characters in length.
    Description string `json:"description"`


    

}

// String returns a JSON representation of the model
func (o *Scheduledtrigger) String() string {
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scheduledtrigger) MarshalJSON() ([]byte, error) {
    type Alias Scheduledtrigger

    if ScheduledtriggerMarshalled {
        return []byte("{}"), nil
    }
    ScheduledtriggerMarshalled = true

    return json.Marshal(&struct {
        
        Name string `json:"name"`
        
        Target Triggertarget `json:"target"`
        
        Version int `json:"version"`
        
        Enabled bool `json:"enabled"`
        
        Schedule Triggerschedule `json:"schedule"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        


        


        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

