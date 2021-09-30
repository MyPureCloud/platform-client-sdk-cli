package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    ScheduleintervalMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type ScheduleintervalDud struct { 
    


    

}

// Scheduleinterval
type Scheduleinterval struct { 
    // Start - The scheduled start time as an ISO-8601 string, i.e yyyy-MM-ddTHH:mm:ss.SSSZ
    Start string `json:"start"`


    // End - The scheduled end time as an ISO-8601 string, i.e. yyyy-MM-ddTHH:mm:ss.SSSZ
    End string `json:"end"`

}

// String returns a JSON representation of the model
func (o *Scheduleinterval) String() string {
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Scheduleinterval) MarshalJSON() ([]byte, error) {
    type Alias Scheduleinterval

    if ScheduleintervalMarshalled {
        return []byte("{}"), nil
    }
    ScheduleintervalMarshalled = true

    return json.Marshal(&struct { 
        Start string `json:"start"`
        
        End string `json:"end"`
        
        *Alias
    }{
        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

