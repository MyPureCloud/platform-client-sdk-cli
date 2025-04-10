package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuagentscheduleshiftrequestMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuagentscheduleshiftrequestDud struct { 
    


    StartDate time.Time `json:"startDate"`


    LengthMinutes int `json:"lengthMinutes"`


    


    


    Schedule Buschedulereference `json:"schedule"`

}

// Buagentscheduleshiftrequest
type Buagentscheduleshiftrequest struct { 
    // Id - The ID of the shift
    Id string `json:"id"`


    


    


    // Activities - The activities associated with this shift
    Activities []Buagentscheduleactivity `json:"activities"`


    // ManuallyEdited - Whether this shift was manually edited. This is only set by clients and is used for rescheduling
    ManuallyEdited bool `json:"manuallyEdited"`


    

}

// String returns a JSON representation of the model
func (o *Buagentscheduleshiftrequest) String() string {
    
     o.Activities = []Buagentscheduleactivity{{}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buagentscheduleshiftrequest) MarshalJSON() ([]byte, error) {
    type Alias Buagentscheduleshiftrequest

    if BuagentscheduleshiftrequestMarshalled {
        return []byte("{}"), nil
    }
    BuagentscheduleshiftrequestMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        
        Activities []Buagentscheduleactivity `json:"activities"`
        
        ManuallyEdited bool `json:"manuallyEdited"`
        *Alias
    }{

        


        


        


        
        Activities: []Buagentscheduleactivity{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

