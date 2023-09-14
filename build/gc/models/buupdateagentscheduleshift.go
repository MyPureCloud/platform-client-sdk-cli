package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuupdateagentscheduleshiftMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuupdateagentscheduleshiftDud struct { 
    


    StartDate time.Time `json:"startDate"`


    LengthMinutes int `json:"lengthMinutes"`


    


    


    Schedule Buschedulereference `json:"schedule"`


    

}

// Buupdateagentscheduleshift
type Buupdateagentscheduleshift struct { 
    // Id - The ID of the shift
    Id string `json:"id"`


    


    


    // Activities - The activities associated with this shift
    Activities []Buagentscheduleactivity `json:"activities"`


    // ManuallyEdited - Whether this shift was manually edited. This is only set by clients and is used for rescheduling
    ManuallyEdited bool `json:"manuallyEdited"`


    


    // Delete - Set to true to delete the shift from the agent's schedule
    Delete bool `json:"delete"`

}

// String returns a JSON representation of the model
func (o *Buupdateagentscheduleshift) String() string {
    
     o.Activities = []Buagentscheduleactivity{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buupdateagentscheduleshift) MarshalJSON() ([]byte, error) {
    type Alias Buupdateagentscheduleshift

    if BuupdateagentscheduleshiftMarshalled {
        return []byte("{}"), nil
    }
    BuupdateagentscheduleshiftMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Activities []Buagentscheduleactivity `json:"activities"`
        
        ManuallyEdited bool `json:"manuallyEdited"`
        
        Delete bool `json:"delete"`
        *Alias
    }{

        


        


        


        
        Activities: []Buagentscheduleactivity{{}},
        


        


        


        

        Alias: (*Alias)(u),
    })
}

