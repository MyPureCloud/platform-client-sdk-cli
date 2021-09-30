package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    TimeallowedMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type TimeallowedDud struct { 
    


    


    

}

// Timeallowed
type Timeallowed struct { 
    // TimeSlots
    TimeSlots []Timeslot `json:"timeSlots"`


    // TimeZoneId
    TimeZoneId string `json:"timeZoneId"`


    // Empty
    Empty bool `json:"empty"`

}

// String returns a JSON representation of the model
func (o *Timeallowed) String() string {
    
    
     o.TimeSlots = []Timeslot{{}} 
    
    
    
    
    
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Timeallowed) MarshalJSON() ([]byte, error) {
    type Alias Timeallowed

    if TimeallowedMarshalled {
        return []byte("{}"), nil
    }
    TimeallowedMarshalled = true

    return json.Marshal(&struct { 
        TimeSlots []Timeslot `json:"timeSlots"`
        
        TimeZoneId string `json:"timeZoneId"`
        
        Empty bool `json:"empty"`
        
        *Alias
    }{
        

        
        TimeSlots: []Timeslot{{}},
        

        

        

        

        

        
        Alias: (*Alias)(u),
    })
}

