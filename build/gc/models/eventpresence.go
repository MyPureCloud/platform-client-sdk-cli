package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    EventpresenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type EventpresenceDud struct { 
    

}

// Eventpresence - A Presence event.
type Eventpresence struct { 
    // VarType - Describes the type of Presence event.
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Eventpresence) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Eventpresence) MarshalJSON() ([]byte, error) {
    type Alias Eventpresence

    if EventpresenceMarshalled {
        return []byte("{}"), nil
    }
    EventpresenceMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

