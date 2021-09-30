package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    WebchatsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type WebchatsettingsDud struct { 
    

}

// Webchatsettings
type Webchatsettings struct { 
    // RequireDeployment
    RequireDeployment bool `json:"requireDeployment"`

}

// String returns a JSON representation of the model
func (o *Webchatsettings) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Webchatsettings) MarshalJSON() ([]byte, error) {
    type Alias Webchatsettings

    if WebchatsettingsMarshalled {
        return []byte("{}"), nil
    }
    WebchatsettingsMarshalled = true

    return json.Marshal(&struct { 
        RequireDeployment bool `json:"requireDeployment"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

