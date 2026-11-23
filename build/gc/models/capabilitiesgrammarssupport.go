package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesgrammarssupportMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesgrammarssupportDud struct { 
    

}

// Capabilitiesgrammarssupport
type Capabilitiesgrammarssupport struct { 
    // Limits
    Limits map[string]int `json:"limits"`

}

// String returns a JSON representation of the model
func (o *Capabilitiesgrammarssupport) String() string {
     o.Limits = map[string]int{"": 0} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiesgrammarssupport) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiesgrammarssupport

    if CapabilitiesgrammarssupportMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesgrammarssupportMarshalled = true

    return json.Marshal(&struct {
        
        Limits map[string]int `json:"limits"`
        *Alias
    }{

        
        Limits: map[string]int{"": 0},
        

        Alias: (*Alias)(u),
    })
}

