package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CallabletimeMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CallabletimeDud struct { 
    


    


    

}

// Callabletime
type Callabletime struct { 
    // TimeSlots - The time intervals for which it is acceptable to place outbound calls.
    TimeSlots []Campaigntimeslot `json:"timeSlots"`


    // Name - The name for the callable time.
    Name string `json:"name"`


    // TimeZoneId - The time zone for the time slots; for example, Africa/Abidjan. Time zones are represented as a string of the zone name as found in the IANA time zone database. For example: UTC, Etc/UTC, or Europe/London
    TimeZoneId string `json:"timeZoneId"`

}

// String returns a JSON representation of the model
func (o *Callabletime) String() string {
     o.TimeSlots = []Campaigntimeslot{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Callabletime) MarshalJSON() ([]byte, error) {
    type Alias Callabletime

    if CallabletimeMarshalled {
        return []byte("{}"), nil
    }
    CallabletimeMarshalled = true

    return json.Marshal(&struct {
        
        TimeSlots []Campaigntimeslot `json:"timeSlots"`
        
        Name string `json:"name"`
        
        TimeZoneId string `json:"timeZoneId"`
        *Alias
    }{

        
        TimeSlots: []Campaigntimeslot{{}},
        


        


        

        Alias: (*Alias)(u),
    })
}

