package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeslotMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeslotDud struct { 
    


    


    

}

// Timeslot
type Timeslot struct { 
    // StartTime - start time in xx:xx:xx.xxx format
    StartTime string `json:"startTime"`


    // StopTime - stop time in xx:xx:xx.xxx format
    StopTime string `json:"stopTime"`


    // Day - Day for this time slot, Monday = 1 ... Sunday = 7
    Day int `json:"day"`

}

// String returns a JSON representation of the model
func (o *Timeslot) String() string {
    
    
    
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeslot) MarshalJSON() ([]byte, error) {
    type Alias Timeslot

    if TimeslotMarshalled {
        return []byte("{}"), nil
    }
    TimeslotMarshalled = true

    return json.Marshal(&struct { 
        StartTime string `json:"startTime"`
        
        StopTime string `json:"stopTime"`
        
        Day int `json:"day"`
        
        *Alias
    }{
        

        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

