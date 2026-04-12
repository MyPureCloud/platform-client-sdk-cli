package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CreatescheduledtriggerrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CreatescheduledtriggerrequestDud struct { 
    


    


    


    


    

}

// Createscheduledtriggerrequest
type Createscheduledtriggerrequest struct { 
    // Target - The target to invoke when the scheduled trigger fires
    Target Triggertarget `json:"target"`


    // Enabled - Boolean indicating if scheduled trigger is enabled
    Enabled bool `json:"enabled"`


    // Name - The name of the scheduled trigger. Can be up to 162 characters in length.
    Name string `json:"name"`


    // Schedule - The schedule configuration for when this trigger should fire
    Schedule Triggerschedule `json:"schedule"`


    // Description - Description of the trigger. Can be up to 512 characters in length.
    Description string `json:"description"`

}

// String returns a JSON representation of the model
func (o *Createscheduledtriggerrequest) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Createscheduledtriggerrequest) MarshalJSON() ([]byte, error) {
    type Alias Createscheduledtriggerrequest

    if CreatescheduledtriggerrequestMarshalled {
        return []byte("{}"), nil
    }
    CreatescheduledtriggerrequestMarshalled = true

    return json.Marshal(&struct {
        
        Target Triggertarget `json:"target"`
        
        Enabled bool `json:"enabled"`
        
        Name string `json:"name"`
        
        Schedule Triggerschedule `json:"schedule"`
        
        Description string `json:"description"`
        *Alias
    }{

        


        


        


        


        

        Alias: (*Alias)(u),
    })
}

