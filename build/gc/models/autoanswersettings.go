package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AutoanswersettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AutoanswersettingsDud struct { 
    

}

// Autoanswersettings
type Autoanswersettings struct { 
    // Settings - Map of conversation media type enabled status.
    Settings map[string]Autoanswersetting `json:"settings"`

}

// String returns a JSON representation of the model
func (o *Autoanswersettings) String() string {
     o.Settings = map[string]Autoanswersetting{"": {}} 

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Autoanswersettings) MarshalJSON() ([]byte, error) {
    type Alias Autoanswersettings

    if AutoanswersettingsMarshalled {
        return []byte("{}"), nil
    }
    AutoanswersettingsMarshalled = true

    return json.Marshal(&struct {
        
        Settings map[string]Autoanswersetting `json:"settings"`
        *Alias
    }{

        
        Settings: map[string]Autoanswersetting{"": {}},
        

        Alias: (*Alias)(u),
    })
}

