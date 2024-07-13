package models
import (
    "time"
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecurrenceendsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecurrenceendsettingsDud struct { 
    


    

}

// Recurrenceendsettings
type Recurrenceendsettings struct { 
    // LastDate - The end date of the recurrence for the activity plan, in ISO-8601 format. Only one of lastDate or noEndDate may be set
    LastDate time.Time `json:"lastDate"`


    // NoEndDate - Whether this activity plan should continue indefinitely. If set to true, lastDate must not be set
    NoEndDate bool `json:"noEndDate"`

}

// String returns a JSON representation of the model
func (o *Recurrenceendsettings) String() string {
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recurrenceendsettings) MarshalJSON() ([]byte, error) {
    type Alias Recurrenceendsettings

    if RecurrenceendsettingsMarshalled {
        return []byte("{}"), nil
    }
    RecurrenceendsettingsMarshalled = true

    return json.Marshal(&struct {
        
        LastDate time.Time `json:"lastDate"`
        
        NoEndDate bool `json:"noEndDate"`
        *Alias
    }{

        


        

        Alias: (*Alias)(u),
    })
}

