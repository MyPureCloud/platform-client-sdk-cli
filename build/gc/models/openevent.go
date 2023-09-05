package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    OpeneventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type OpeneventDud struct { 
    

}

// Openevent - Message event element.
type Openevent struct { 
    // EventType - Type of this event element
    EventType string `json:"eventType"`

}

// String returns a JSON representation of the model
func (o *Openevent) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Openevent) MarshalJSON() ([]byte, error) {
    type Alias Openevent

    if OpeneventMarshalled {
        return []byte("{}"), nil
    }
    OpeneventMarshalled = true

    return json.Marshal(&struct {
        
        EventType string `json:"eventType"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

