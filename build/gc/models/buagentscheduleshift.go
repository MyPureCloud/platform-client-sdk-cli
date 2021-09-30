package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    BuagentscheduleshiftMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type BuagentscheduleshiftDud struct { 
    Id string `json:"id"`


    


    


    


    


    Schedule Buschedulereference `json:"schedule"`

}

// Buagentscheduleshift
type Buagentscheduleshift struct { 
    


    // StartDate - The start date of this shift. Date time is represented as an ISO-8601 string. For example: yyyy-MM-ddTHH:mm:ss[.mmm]Z
    StartDate time.Time `json:"startDate"`


    // LengthMinutes - The length of this shift in minutes
    LengthMinutes int `json:"lengthMinutes"`


    // Activities - The activities associated with this shift
    Activities []Buagentscheduleactivity `json:"activities"`


    // ManuallyEdited - Whether this shift was manually edited. This is only set by clients and is used for rescheduling
    ManuallyEdited bool `json:"manuallyEdited"`


    

}

// String returns a JSON representation of the model
func (o *Buagentscheduleshift) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
     o.Activities = []Buagentscheduleactivity{{}} 
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Buagentscheduleshift) MarshalJSON() ([]byte, error) {
    type Alias Buagentscheduleshift

    if BuagentscheduleshiftMarshalled {
        return []byte("{}"), nil
    }
    BuagentscheduleshiftMarshalled = true

    return json.Marshal(&struct { 
        
        
        StartDate time.Time `json:"startDate"`
        
        LengthMinutes int `json:"lengthMinutes"`
        
        Activities []Buagentscheduleactivity `json:"activities"`
        
        ManuallyEdited bool `json:"manuallyEdited"`
        
        
        
        *Alias
    }{
        

        

        

        

        

        

        

        
        Activities: []Buagentscheduleactivity{{}},
        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

