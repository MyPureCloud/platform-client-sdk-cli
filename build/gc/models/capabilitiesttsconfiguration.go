package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesttsconfigurationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesttsconfigurationDud struct { 
    

}

// Capabilitiesttsconfiguration
type Capabilitiesttsconfiguration struct { 
    // Engines
    Engines map[string]Capabilitiesttsengine `json:"engines"`

}

// String returns a JSON representation of the model
func (o *Capabilitiesttsconfiguration) String() string {
     o.Engines = map[string]Capabilitiesttsengine{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiesttsconfiguration) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiesttsconfiguration

    if CapabilitiesttsconfigurationMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesttsconfigurationMarshalled = true

    return json.Marshal(&struct {
        
        Engines map[string]Capabilitiesttsengine `json:"engines"`
        *Alias
    }{

        
        Engines: map[string]Capabilitiesttsengine{"": {}},
        

        Alias: (*Alias)(u),
    })
}

