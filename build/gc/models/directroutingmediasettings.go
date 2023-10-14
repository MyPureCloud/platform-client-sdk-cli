package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    DirectroutingmediasettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type DirectroutingmediasettingsDud struct { 
    

}

// Directroutingmediasettings
type Directroutingmediasettings struct { 
    // UseAgentAddressOutbound - Toggle that enables using an agent's Direct Routing address outbound on behalf of queue for this media type.
    UseAgentAddressOutbound bool `json:"useAgentAddressOutbound"`

}

// String returns a JSON representation of the model
func (o *Directroutingmediasettings) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Directroutingmediasettings) MarshalJSON() ([]byte, error) {
    type Alias Directroutingmediasettings

    if DirectroutingmediasettingsMarshalled {
        return []byte("{}"), nil
    }
    DirectroutingmediasettingsMarshalled = true

    return json.Marshal(&struct {
        
        UseAgentAddressOutbound bool `json:"useAgentAddressOutbound"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

