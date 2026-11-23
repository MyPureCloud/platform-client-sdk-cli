package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiessttconfigurationMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiessttconfigurationDud struct { 
    


    

}

// Capabilitiessttconfiguration
type Capabilitiessttconfiguration struct { 
    // SttEngines
    SttEngines map[string]Capabilitiessttengine `json:"sttEngines"`


    // SttSettings
    SttSettings Capabilitiessttsettings `json:"sttSettings"`

}

// String returns a JSON representation of the model
func (o *Capabilitiessttconfiguration) String() string {
     o.SttEngines = map[string]Capabilitiessttengine{"": {}} 
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiessttconfiguration) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiessttconfiguration

    if CapabilitiessttconfigurationMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiessttconfigurationMarshalled = true

    return json.Marshal(&struct {
        
        SttEngines map[string]Capabilitiessttengine `json:"sttEngines"`
        
        SttSettings Capabilitiessttsettings `json:"sttSettings"`
        *Alias
    }{

        
        SttEngines: map[string]Capabilitiessttengine{"": {}},
        


        

        Alias: (*Alias)(u),
    })
}

