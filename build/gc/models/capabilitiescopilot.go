package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiescopilotMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiescopilotDud struct { 
    Enabled bool `json:"enabled"`

}

// Capabilitiescopilot
type Capabilitiescopilot struct { 
    

}

// String returns a JSON representation of the model
func (o *Capabilitiescopilot) String() string {

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiescopilot) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiescopilot

    if CapabilitiescopilotMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiescopilotMarshalled = true

    return json.Marshal(&struct {
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

