package models
import (
    "encoding/json"
    "strconv"
    "strings"
)

var (
    CrossplatformpolicyupdateMarshalled = false
)

// This struct is here to use the useless readonly properties so that their required imports don't throw an unused error (time, etc.)
type CrossplatformpolicyupdateDud struct { 
    

}

// Crossplatformpolicyupdate
type Crossplatformpolicyupdate struct { 
    // Enabled
    Enabled bool `json:"enabled"`

}

// String returns a JSON representation of the model
func (o *Crossplatformpolicyupdate) String() string {
    
    
    
    
    

    j, _ := json.Marshal(o)
    str, _ := strconv.Unquote(strings.Replace(strconv.Quote(string(j)), `\\u`, `\u`, -1))

    return str
}

func (u *Crossplatformpolicyupdate) MarshalJSON() ([]byte, error) {
    type Alias Crossplatformpolicyupdate

    if CrossplatformpolicyupdateMarshalled {
        return []byte("{}"), nil
    }
    CrossplatformpolicyupdateMarshalled = true

    return json.Marshal(&struct { 
        Enabled bool `json:"enabled"`
        
        *Alias
    }{
        

        

        
        Alias: (*Alias)(u),
    })
}

