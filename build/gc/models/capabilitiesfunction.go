package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesfunctionMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesfunctionDud struct { 
    

}

// Capabilitiesfunction
type Capabilitiesfunction struct { 
    // Overloads
    Overloads map[string]Capabilitiesflowtypesavailability `json:"overloads"`

}

// String returns a JSON representation of the model
func (o *Capabilitiesfunction) String() string {
     o.Overloads = map[string]Capabilitiesflowtypesavailability{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiesfunction) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiesfunction

    if CapabilitiesfunctionMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesfunctionMarshalled = true

    return json.Marshal(&struct {
        
        Overloads map[string]Capabilitiesflowtypesavailability `json:"overloads"`
        *Alias
    }{

        
        Overloads: map[string]Capabilitiesflowtypesavailability{"": {}},
        

        Alias: (*Alias)(u),
    })
}

