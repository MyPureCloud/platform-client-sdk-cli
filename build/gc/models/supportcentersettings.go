package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    SupportcentersettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type SupportcentersettingsDud struct { 
    

}

// Supportcentersettings - Settings concerning support center
type Supportcentersettings struct { 
    // Enabled - Whether or not support center is enabled
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Supportcentersettings) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Supportcentersettings) MarshalJSON() ([]byte, error) {
    type Alias Supportcentersettings

    if SupportcentersettingsMarshalled {
        return []byte("{}"), nil
    }
    SupportcentersettingsMarshalled = true

    return json.Marshal(&struct { 
        Enabled bool `json:"enabled"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}
