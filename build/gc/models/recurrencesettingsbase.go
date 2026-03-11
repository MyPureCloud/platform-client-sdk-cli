package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    RecurrencesettingsbaseMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type RecurrencesettingsbaseDud struct { 
    

}

// Recurrencesettingsbase
type Recurrencesettingsbase struct { 
    // EndAfter - Settings controlling when to end the recurrence for the activity plan
    EndAfter Recurrenceendsettings `json:"endAfter"`

}

// String returns a JSON representation of the model
func (o *Recurrencesettingsbase) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Recurrencesettingsbase) MarshalJSON() ([]byte, error) {
    type Alias Recurrencesettingsbase

    if RecurrencesettingsbaseMarshalled {
        return []byte("{}"), nil
    }
    RecurrencesettingsbaseMarshalled = true

    return json.Marshal(&struct {
        
        EndAfter Recurrenceendsettings `json:"endAfter"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

