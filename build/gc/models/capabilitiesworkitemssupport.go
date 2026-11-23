package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesworkitemssupportMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesworkitemssupportDud struct { 
    

}

// Capabilitiesworkitemssupport
type Capabilitiesworkitemssupport struct { 
    // Events
    Events []Capabilitiesworkitemevent `json:"events"`

}

// String returns a JSON representation of the model
func (o *Capabilitiesworkitemssupport) String() string {
     o.Events = []Capabilitiesworkitemevent{{}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiesworkitemssupport) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiesworkitemssupport

    if CapabilitiesworkitemssupportMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesworkitemssupportMarshalled = true

    return json.Marshal(&struct {
        
        Events []Capabilitiesworkitemevent `json:"events"`
        *Alias
    }{

        
        Events: []Capabilitiesworkitemevent{{}},
        

        Alias: (*Alias)(u),
    })
}

