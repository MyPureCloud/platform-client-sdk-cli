package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CapabilitiesintentdetectionsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CapabilitiesintentdetectionsettingsDud struct { 
    

}

// Capabilitiesintentdetectionsettings
type Capabilitiesintentdetectionsettings struct { 
    // VarType
    VarType string `json:"type"`

}

// String returns a JSON representation of the model
func (o *Capabilitiesintentdetectionsettings) String() string {
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Capabilitiesintentdetectionsettings) MarshalJSON() ([]byte, error) {
    type Alias Capabilitiesintentdetectionsettings

    if CapabilitiesintentdetectionsettingsMarshalled {
        return []byte("{}"), nil
    }
    CapabilitiesintentdetectionsettingsMarshalled = true

    return json.Marshal(&struct {
        
        VarType string `json:"type"`
        *Alias
    }{

        

        Alias: (*Alias)(u),
    })
}

