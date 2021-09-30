package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WeekschedulelistresponseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WeekschedulelistresponseDud struct { 
    

}

// Weekschedulelistresponse - Week schedule list
type Weekschedulelistresponse struct { 
    // Entities
    Entities []Weekschedulelistitemresponse `json:"entities"`

}

// String returns a JSON representation of the model
func (o *Weekschedulelistresponse) String() string {
    
    
     o.Entities = []Weekschedulelistitemresponse{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Weekschedulelistresponse) MarshalJSON() ([]byte, error) {
    type Alias Weekschedulelistresponse

    if WeekschedulelistresponseMarshalled {
        return []byte("{}"), nil
    }
    WeekschedulelistresponseMarshalled = true

    return json.Marshal(&struct { 
        Entities []Weekschedulelistitemresponse `json:"entities"`
        
        *Alias
    }{
        

        
        Entities: []Weekschedulelistitemresponse{{}},
        

        
        Alias: (*Alias)(u),
    })
}

