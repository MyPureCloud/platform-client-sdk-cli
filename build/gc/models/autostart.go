package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AutostartMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AutostartDud struct { 
    

}

// Autostart
type Autostart struct { 
    // Enabled - whether or not auto start is enabled
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Autostart) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Autostart) MarshalJSON() ([]byte, error) {
    type Alias Autostart

    if AutostartMarshalled {
        return []byte("{}"), nil
    }
    AutostartMarshalled = true

    return json.Marshal(&struct { 
        Enabled bool `json:"enabled"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

