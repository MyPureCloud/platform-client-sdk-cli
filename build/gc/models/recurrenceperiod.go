package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecurrenceperiodMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecurrenceperiodDud struct { 
    


    

}

// Recurrenceperiod
type Recurrenceperiod struct { 
    // Magnitude - The period of the activity plan in granularity units
    Magnitude int `json:"magnitude"`


    // Granularity - The granularity unit to interpret the period of this activity plan
    Granularity string `json:"granularity"`

}

// String returns a JSON representation of the model
func (o *Recurrenceperiod) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recurrenceperiod) MarshalJSON() ([]byte, error) {
    type Alias Recurrenceperiod

    if RecurrenceperiodMarshalled {
        return []byte("{}"), nil
    }
    RecurrenceperiodMarshalled = true

    return json.Marshal(&struct {
        
        Magnitude int `json:"magnitude"`
        
        Granularity string `json:"granularity"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

