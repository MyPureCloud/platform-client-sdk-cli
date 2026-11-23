package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesworkitemeventMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesworkitemeventDud struct { 
    

}

// Capabilitiesworkitemevent
type Capabilitiesworkitemevent struct { 
    // Id
    Id string `json:"id"`

}

// String returns a JSON representation of the model
func (o *Capabilitiesworkitemevent) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiesworkitemevent) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiesworkitemevent

    if CapabilitiesworkitemeventMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesworkitemeventMarshalled = true

    return json.Marshal(&struct {
        
        Id string `json:"id"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

