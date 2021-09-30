package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    AutomatictimezonemappingsettingsMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type AutomatictimezonemappingsettingsDud struct { 
    

}

// Automatictimezonemappingsettings
type Automatictimezonemappingsettings struct { 
    // CallableWindows - The time intervals to use for automatic time zone mapping.
    CallableWindows []Callablewindow `json:"callableWindows"`

}

// String returns a JSON representation of the model
func (o *Automatictimezonemappingsettings) String() string {
    
    
     o.CallableWindows = []Callablewindow{{}} 
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Automatictimezonemappingsettings) MarshalJSON() ([]byte, error) {
    type Alias Automatictimezonemappingsettings

    if AutomatictimezonemappingsettingsMarshalled {
        return []byte("{}"), nil
    }
    AutomatictimezonemappingsettingsMarshalled = true

    return json.Marshal(&struct { 
        CallableWindows []Callablewindow `json:"callableWindows"`
        
        *Alias
    }{
        

        
        CallableWindows: []Callablewindow{{}},
        

        
        Alias: (*Alias)(u),
    })
}

