package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebmessagingeventpresenceMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebmessagingeventpresenceDud struct { 
    

}

// Webmessagingeventpresence - A Presence event.
type Webmessagingeventpresence struct { 
    // VarType - Describes the type of Presence event.
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Webmessagingeventpresence) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webmessagingeventpresence) MarshalJSON() ([]byte, error) {
    type Alias Webmessagingeventpresence

    if WebmessagingeventpresenceMarshalled {
        return []byte("{}"), nil
    }
    WebmessagingeventpresenceMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

